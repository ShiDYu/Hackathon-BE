package usecase

import (
	"api/dao"
	"api/model"
)

func GetTweets() ([]model.Tweet, error) {
	// ツイート情報をデータベースから取得する処理
	return dao.GetTweets()
}

func CreateTweet(tweet model.Tweet) error {
	return dao.CreateTweet(tweet)
}

func UpdateTweet(id int, content string) error {
	return dao.UpdateTweetContent(id, content)
}

func GetTodayTweetCount(UserID string) (int, error) {
	return dao.GetTodayTweetCount(UserID)
}
