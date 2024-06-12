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
