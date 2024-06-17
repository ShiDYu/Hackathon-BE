package dao

import (
	"api/model"
	"log"
)

func GetLikes(postID, userID string) (model.PostResponse, error) {

	var post model.LikedPost
	err := db.QueryRow("SELECT id, like_count FROM tweets WHERE id = ?", postID).Scan(&post.ID, &post.LikeCount)
	if err != nil {
		return model.PostResponse{}, nil
	}

	rows, err := db.Query("SELECT user_id FROM likes WHERE tweet_id = ?", postID)
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
	post.Likes = likes

	response := model.PostResponse{
		LikeCount:   post.LikeCount,
		LikedByUser: likedByUser,
	}

	return response, nil
}

func LikePost(like model.Like) (model.LikedPost, error) {

	tx, err := db.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v\n", err)
		return model.LikedPost{}, err
	}
	log.Println("Transaction started")

	_, err = tx.Exec("INSERT INTO likes (tweet_id, user_id) VALUES (?, ?) ON DUPLICATE KEY UPDATE tweet_id = tweet_id", like.PostID, like.UserID)
	if err != nil {
		log.Printf("Error inserting like: %v\n", err)
		tx.Rollback()
		return model.LikedPost{}, err
	}
	log.Println("Inserted like or duplicate detected")

	_, err = tx.Exec("UPDATE tweets SET like_count = like_count + 1 WHERE id = ?", like.PostID)
	if err != nil {
		log.Printf("Error updating like count: %v\n", err)
		tx.Rollback()
		return model.LikedPost{}, err
	}
	log.Println("Updated like count")

	err = tx.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v\n", err)
		return model.LikedPost{}, err
	}
	log.Println("Transaction committed")

	var post model.LikedPost
	err = db.QueryRow("SELECT id, like_count FROM tweets WHERE id = ?", like.PostID).Scan(&post.ID, &post.LikeCount)
	if err != nil {
		log.Printf("Error fetching post: %v\n", err)
		return model.LikedPost{}, err
	}
	log.Printf("Fetched post: %+v\n", post)

	log.Println("LikePost function completed successfully")
	return post, nil
}

//いいねカウントがバグってるのであとで直す 一回クリックするとカウントの値から取得できている
//likeカウントがマイナスにならないようにバリデーションをかける

func UnlikePost(like model.Like) (model.LikedPost, error) {

	tx, err := db.Begin()
	if err != nil {
		return model.LikedPost{}, err
	}

	_, err = tx.Exec("DELETE FROM likes WHERE tweet_id = ? AND user_id = ?", like.PostID, like.UserID)
	if err != nil {
		tx.Rollback()
		return model.LikedPost{}, err
	}

	_, err = tx.Exec("UPDATE tweets SET like_count = like_count - 1 WHERE id = ?", like.PostID)
	if err != nil {
		tx.Rollback()
		return model.LikedPost{}, err
	}

	err = tx.Commit()
	if err != nil {
		return model.LikedPost{}, err
	}

	var post model.LikedPost
	err = db.QueryRow("SELECT id, like_count FROM tweets WHERE id = ?", like.PostID).Scan(&post.ID, &post.LikeCount)
	if err != nil {
		return model.LikedPost{}, err
	}

	//json.NewEncoder(w).Encode(map[string]int{"likeCount": post.LikeCount})
	return post, nil
}
