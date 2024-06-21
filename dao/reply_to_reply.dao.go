package dao

import (
	"api/model"
	"log"
)

func CreateReplyToReply(reply model.ReplyToReply) error {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return err
	}

	query := `INSERT INTO replies (reply_id, uid, content) VALUES (?, ?, ?)`
	_, err = tx.Exec(query, reply.ReplyId, reply.Uid, reply.Content)
	if err != nil {
		tx.Rollback()
		log.Printf("Error executing query: %v", err)
		return err
	}
	err = tx.Commit() // Add commit to finalize the transaction
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
		return err
	}
	return nil
}

func GetRepliesToReply(replyID int) ([]model.ReplyToReply, error) {
	query := `SELECT id, reply_id, uid, content, created_at FROM replies WHERE reply_id = ?`
	rows, err := db.Query(query, replyID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var replies []model.ReplyToReply
	for rows.Next() {
		var reply model.ReplyToReply
		if err := rows.Scan(&reply.Id, &reply.ReplyId, &reply.Uid, &reply.Content, &reply.CreatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
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

func LikeReply(replyID int, uid string) (model.LikedReply, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Println("Transaction started")

	result, err := tx.Exec("INSERT INTO likeReply (reply_id, uid) VALUES (?, ?) ON DUPLICATE KEY UPDATE reply_id = reply_id", replyID, uid)
	if err != nil {
		log.Printf("Error inserting like: %v\n", err)
		tx.Rollback()
		return model.LikedReply{}, err
	}
	log.Println("Inserted like or duplicate detected")

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return model.LikedReply{}, err
	}

	_, err = tx.Exec("UPDATE replies SET like_count = like_count + ? WHERE id = ?", rowsAffected, replyID)
	if err != nil {
		log.Printf("Error updating like count: %v\n", err)
		tx.Rollback()
		return model.LikedReply{}, err
	}
	log.Println("Updated like count")

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Println("Transaction committed")

	var reply model.LikedReply
	err = db.QueryRow("SELECT id, like_count FROM replies WHERE id = ?", replyID).Scan(&reply.ID, &reply.LikeCount)
	if err != nil {
		log.Printf("Error fetching reply: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Printf("Fetched reply: %+v\n", reply)

	log.Println("LikeReply function completed successfully")
	return reply, nil
}

func UnLikeReply(replyID int, uid string) (model.LikedReply, error) {
	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Println("Transaction started")

	query := `DELETE FROM likeReply WHERE reply_id=? AND uid=?`
	result, err := tx.Exec(query, replyID, uid)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		tx.Rollback()
		return model.LikedReply{}, err
	}
	log.Println("Deleted like")

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return model.LikedReply{}, err
	}

	query = `UPDATE replies SET like_count = like_count - ? WHERE id = ?`
	_, err = tx.Exec(query, rowsAffected, replyID)
	if err != nil {
		log.Printf("Error updating like count: %v\n", err)
		tx.Rollback()
		return model.LikedReply{}, err
	}
	log.Println("Updated like count")

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Println("Transaction committed")

	var reply model.LikedReply
	err = db.QueryRow("SELECT id, like_count FROM replies WHERE id = ?", replyID).Scan(&reply.ID, &reply.LikeCount)
	if err != nil {
		log.Printf("Error fetching reply: %v\n", err)
		return model.LikedReply{}, err
	}
	log.Printf("Fetched reply: %+v\n", reply)

	log.Println("UnLikeReply function completed successfully")
	return reply, nil
}
func GetReplycount(replyID int) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM replies WHERE reply_id = ?`
	err := db.QueryRow(query, replyID).Scan(&count)
	return count, err
}

//アルゴリズム違う

func GetReplyLikes(replyID int, userID string) (model.PostResponse, error) {

	var reply model.ReplyToReply
	err := db.QueryRow("SELECT id, like_count FROM replies WHERE id = ?", replyID).Scan(&reply.Id, &reply.LikeCount)
	if err != nil {
		return model.PostResponse{}, nil
	}

	rows, err := db.Query("SELECT uid FROM likeReply WHERE reply_id = ?", replyID)
	if err != nil {
		return model.PostResponse{}, nil
	}
	defer rows.Close()

	var likes []string
	var likedByUser bool
	for rows.Next() {
		var uid string
		err := rows.Scan(&uid)
		if err != nil {
			return model.PostResponse{}, nil
		}
		likes = append(likes, uid)
		if uid == userID {
			likedByUser = true
		}
	}

	response := model.PostResponse{
		LikeCount:   reply.LikeCount,
		LikedByUser: likedByUser,
	}

	return response, nil
}
