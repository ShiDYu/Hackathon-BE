package dao

import (
	"api/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func InitDB() {
	//デプロイする時はここの部分を毎回コメントアウトする
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// ①-1』
	//デプロイする時はここまでを毎回コメントアウトする
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlUserPwd := os.Getenv("MYSQL_USER_PWD")
	mysqlDatabase := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")

	// ①-2
	_db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", mysqlUser, mysqlUserPwd, mysqlHost, mysqlDatabase))
	if err != nil {
		log.Fatalf("fail: sql.Open, %v\n", err)
	}
	// ①-3
	if err := _db.Ping(); err != nil {
		log.Fatalf("fail: _db.Ping, %v\n", err)
	}
	db = _db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func GetUserByName(name string) ([]model.UserGet, error) {
	rows, err := db.Query("SELECT id, name, age FROM user WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.UserGet, 0)
	for rows.Next() {
		var u model.UserGet
		if err := rows.Scan(&u.Id, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func InsertUser(user model.UserRegister) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO users (id) VALUES (?) ON DUPLICATE KEY UPDATE id = id", user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func UpdateUserProfile(user model.UserRegister) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET nickname = ?, bio = ? WHERE id = ?", user.Nickname, user.Bio, user.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func SetAvatar(avatar model.AvatarRequest) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE users SET avatar = ? WHERE id = ?", avatar.AvatarURL, avatar.UserId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func GetAvatar(userId string) (string, error) {
	var avatarURL string
	err := db.QueryRow("SELECT avatar FROM users WHERE id = ?", userId).Scan(&avatarURL)
	if err != nil {
		return "", err
	}
	return avatarURL, nil
}

func GetProfile(userId string) (model.UserProfile, error) {
	var profile model.UserProfile
	err := db.QueryRow("SELECT id, nickname, bio, avatar FROM users WHERE id = ?", userId).Scan(&profile.ID, &profile.Nickname, &profile.Bio, &profile.Avatar)
	if err != nil {
		return model.UserProfile{}, err //ポイント型で格納しておけば、空配列ではなく、nilを返す事ができる
	}
	return profile, nil
}
