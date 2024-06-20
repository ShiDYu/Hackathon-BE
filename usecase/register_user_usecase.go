package usecase

import (
	"api/dao"
	"api/model"
)

func RegisterUser(user model.UserRegister) error {
	return dao.InsertUser(user)
}

func UpdateUserProfile(user model.UserRegister) error {
	return dao.UpdateUserProfile(user)
}

func SetAvatar(avatar model.AvatarRequest) error {
	return dao.SetAvatar(avatar)
}

func GetAvatar(userId string) (string, error) {
	return dao.GetAvatar(userId)
}

func GetProfile(userId string) (model.UserProfile, error) {
	return dao.GetProfile(userId)
}
