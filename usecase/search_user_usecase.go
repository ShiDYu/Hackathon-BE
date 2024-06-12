package usecase

import (
	"api/dao"
	"api/model"
)

func SearchUser(name string) ([]model.UserGet, error) {
	return dao.GetUserByName(name)
}
