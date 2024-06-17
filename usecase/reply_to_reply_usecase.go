package usecase

import (
	"api/dao"
	"api/model"
)

func CreateReplyToReply(reply model.ReplyToReply) error {
	return dao.CreateReplyToReply(reply)
}

func GetRepliesTOReply(replyID int) ([]model.ReplyToReply, error) {
	return dao.GetRepliesToReply(replyID)
}

func GetReplyLikes(replyID int, UserID string) (model.PostResponse, error) {
	return dao.GetReplyLikes(replyID, UserID)
}

func ReplyLike(replyID int, Uid string) (model.LikedReply, error) {
	return dao.LikeReply(replyID, Uid)
}

func UnReplyLike(replyID int, Uid string) (model.LikedReply, error) {
	return dao.UnLikeReply(replyID, Uid)
}

func GetReplyToReplyCount(replyID int) (int, error) {
	return dao.GetReplycount(replyID)
}
