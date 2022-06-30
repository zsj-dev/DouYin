package service

import (
	"github.com/zsj-dev/DouYin/basic-service/conf"
	"github.com/zsj-dev/DouYin/database/model"
)

type IUserService interface {
	Create(user *model.User) error
}
type UserService struct{}

func NewUserService() IUserService {
	return UserService{}
}

func (u UserService) Create(user *model.User) error {
	if err := conf.MySQL.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserService) FindByUsername(user *model.User, username string) error {
	if err := conf.MySQL.Model(&model.User{}).
		Where("username = ?", username).
		First(&user).Error; err != nil {
		return err
	}
	return nil
}
