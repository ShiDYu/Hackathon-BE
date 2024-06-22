package dao

import (
	"api/model"
	"log"
	"strconv"
	"time"
)

func GetTweets() ([]model.Tweet, error) {
	query := `SELECT tweets.id, tweets.uid, tweets.content, tweets.created_at, users.nickname, users.avatar FROM tweets JOIN users ON tweets.uid = users.id ORDER BY tweets.created_at DESC`
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
		var avatarURL string
		if err := rows.Scan(&tweet.Id, &tweet.Uid, &tweet.Content, &tweet.Date, &nickname, &avatarURL); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		tweet.Nickname = nickname
		tweet.AvatarURL = avatarURL // ツイートにユーザーネームを追加
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

func UpdateTweetContent(id int, content string) error {

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 内容を更新
	_, err = tx.Exec("UPDATE tweets SET content = ? WHERE id = ?", content, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func GetTodayTweetCount(userID string) (int, error) {

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return 99, err
	}
	today := time.Now().In(loc).Format("2006-01-02")
	log.Printf(today)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM tweets WHERE uid = ? AND DATE(created_at) = ?", userID, today).Scan(&count)
	if err != nil {
		return 55, err
	}
	log.Printf(strconv.Itoa(count))

	return count, nil
}
