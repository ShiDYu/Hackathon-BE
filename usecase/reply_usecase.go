package usecase

import (
	"api/dao"
	"api/model"
)

func CreateReply(reply model.Reply) error {
	return dao.CreateReply(reply)
}

func GetReplies(tweetID int) ([]model.Reply, error) {
	return dao.GetRepliesByTweetID(tweetID)
}

func GetReplyCount(tweetID int) (int, error) {
	return dao.GetReplyCountByTweetID(tweetID)
}

func GetRepliedTweet(tweetID int) (model.Tweet, error) {
	return dao.GetRepliedTweet(tweetID)
}
