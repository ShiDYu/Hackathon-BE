package dao

import (
	"api/model"
	"log"
)

func CreateReply(reply model.Reply) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO replies (tweet_id, uid, content) VALUES (?, ?, ?)",
		reply.TweetId, reply.Uid, reply.Content)
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

//リプライを送るとしっかりDBに保存された。

func GetRepliesByTweetID(tweetID int) ([]model.Reply, error) {
	query := `SELECT id, tweet_id, uid, content, created_at FROM replies WHERE tweet_id = ?`
	rows, err := db.Query(query, tweetID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var replies []model.Reply
	for rows.Next() {
		var reply model.Reply
		if err := rows.Scan(&reply.Id, &reply.TweetId, &reply.Uid, &reply.Content, &reply.Date); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		replies = append(replies, reply)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error with rows: %v", err)
		return nil, err
	}

	return replies, nil
}

func GetReplyCountByTweetID(tweetID int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM replies WHERE tweet_id = ?`
	err := db.QueryRow(query, tweetID).Scan(&count)
	return count, err
}

func GetRepliedTweet(tweetID int) (model.Tweet, error) {

	var tweet model.Tweet
	err := db.QueryRow("SELECT id, uid, content, created_at FROM tweets WHERE id=?", tweetID).Scan(&tweet.Id, &tweet.Uid, &tweet.Content, &tweet.Date)
	if err != nil {
		return model.Tweet{}, err
	}

	return tweet, nil
}
