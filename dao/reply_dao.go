package dao

import (
	"api/model"
	"database/sql"
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

// リプライを送るとしっかりDBに保存された。
func getNicknameByUID(uid string) (string, error) {
	var nickname string
	query := `SELECT nickname FROM users WHERE id = ?`
	err := db.QueryRow(query, uid).Scan(&nickname)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No nickname found for uid: %d", uid)
			return "", nil
		}
		log.Printf("Error executing query: %v", err)
		return "", err
	}
	return nickname, nil
}

func getAvatarByUID(uid string) (string, error) {
	var avatarURL string
	query := `SELECT avatar FROM users WHERE id = ?`
	err := db.QueryRow(query, uid).Scan(&avatarURL)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No avatarURL found for uid: %d", uid)
			return "", nil
		}
		log.Printf("Error executing query: %v", err)
		return "", err
	}
	return avatarURL, nil
}

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

		// ニックネームを取得
		nickname, err := getNicknameByUID(reply.Uid)
		if err != nil {
			log.Printf("Error getting nickname: %v", err)
			return nil, err
		}
		avatarURL, err := getAvatarByUID(reply.Uid)
		log.Printf("AvatarURL: %v", avatarURL)
		if err != nil {
			log.Printf("Error getting avatarURL: %v", err)
			return nil, err
		}
		reply.Nickname = nickname
		reply.AvatarURL = avatarURL

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

func GenerateReply(ReplyContent string, TweetId int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO replies (tweet_id, uid, content) VALUES (?, ?, ?)",
		TweetId, 1, ReplyContent)
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
