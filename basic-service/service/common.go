package service

import (
	"github.com/zsj-dev/DouYin/basic-service/conf"
	"github.com/zsj-dev/DouYin/database/model"
	"gorm.io/gorm"
)

func GetUser(userId int64) (user *model.User, err error) {
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", userId).
		First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func IsFav(UserID int64, videoId int64) (isFav bool, err error) {
	var count int64
	if err := conf.MySQL.Model(&model.Favorite{}).
		Where("user_id = ? AND video_id = ?", UserID, videoId).
		Count(&count).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, err
}

func IsFollow(userId int64, seeId int64) (isFav bool, err error) {
	var count int64
	err = conf.MySQL.Model(&model.Relation{}).
		Where("user_id = ? AND follow_id = ?", userId, seeId).
		Count(&count).Error

	if err == nil && count > 0 {
		return true, err
	}

	return false, err
}
