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
	//デプロイする時はここの部分を毎回コメントアウトする
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

func GetTweets() ([]model.Tweet, error) {
	query := `SELECT tweets.id, tweets.uid, tweets.content, tweets.created_at, users.nickname FROM tweets JOIN users ON tweets.uid = users.id ORDER BY tweets.created_at DESC`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tweets []model.Tweet
	for rows.Next() {
		var tweet model.Tweet
		var nickname string
		if err := rows.Scan(&tweet.Id, &tweet.Uid, &tweet.Content, &tweet.Date, &nickname); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		tweet.Nickname = nickname // ツイートにユーザーネームを追加
		tweets = append(tweets, tweet)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return tweets, nil
}

func CreateTweet(tweet model.Tweet) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO tweets (uid, content) VALUES (?, ?)", tweet.Uid, tweet.Content)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
