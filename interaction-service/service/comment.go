package service

import (
	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/interaction-service/conf"
	"gorm.io/gorm"
)

type ICommentService interface {
	Delete(com *model.Comment) error
	Create(com *model.Comment) error
	UpdateAdd(com *model.Comment) error
	UpdateSub(com *model.Comment) error
	List(videoID int64) (list []model.Comment, err error)
}
type CommentService struct{}

func NewCommentService() ICommentService {
	return CommentService{}
}

func (u CommentService) Delete(com *model.Comment) error {
	if err := conf.MySQL.
		Where("id = ?", com.ID).
		Delete(&model.Comment{}).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) Create(com *model.Comment) error {
	if err := conf.MySQL.Create(&com).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) UpdateAdd(com *model.Comment) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ?", com.VideoId).
		UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) UpdateSub(com *model.Comment) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ? AND comment_count > ?", com.VideoId, 0).
		UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u CommentService) List(videoID int64) (list []model.Comment, err error) {
	if err := conf.MySQL.Model(&model.Comment{}).
		Where("video_id = ?", videoID).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, err
}
