package usecase

import (
	"api/dao"
	"api/model"
	_ "api/model"
)

func GetLikes(postID string, userID string) (model.PostResponse, error) {
	return dao.GetLikes(postID, userID)
}

func LikePost(like model.Like) (model.LikedPost, error) {
	return dao.LikePost(like)
}

func UnlikePost(like model.Like) (model.LikedPost, error) {
	return dao.UnlikePost(like)
}
