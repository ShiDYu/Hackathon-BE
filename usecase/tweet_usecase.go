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
