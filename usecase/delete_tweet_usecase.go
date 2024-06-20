package usecase

import (
	"api/dao"
)

func DeleteTweet(tweetID int) error {
	return dao.DeleteTweet(tweetID)
}

func DeleteReply(replyID int) error {
	return dao.DeleteReply(replyID)
}
