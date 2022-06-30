package service

import (
	"github.com/zsj-dev/DouYin/basic-service/conf"
	"github.com/zsj-dev/DouYin/database/model"
	"gorm.io/gorm"
)

type IPublishService interface {
	Create(m *model.Video) (id int64, err error)
	List(userID int64) (list []model.Video, err error)
}
type PublishService struct{}

func NewPublishService() IPublishService {
	return PublishService{}
}

func (u PublishService) Create(m *model.Video) (id int64, err error) {
	if err := conf.MySQL.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (u PublishService) List(userID int64) (list []model.Video, err error) {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("author_id = ?", userID).
		Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, nil
		}
		return nil, err
	}

	return list, err
}
