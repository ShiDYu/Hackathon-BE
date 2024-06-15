package dao

import (
	"api/model"
	"log"
)

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
