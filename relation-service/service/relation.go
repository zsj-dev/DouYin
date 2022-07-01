package service

import (
	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/relation-service/conf"
	"gorm.io/gorm"
)

type IRelationService interface {
	AddFollow(rel *model.Relation) error
	UnFollow(rel *model.Relation) error
	GetFollowByID(userID int64) (followIds []int64, err error)
	UpdateSub(rel *model.Relation) error
	UpdateAdd(rel *model.Relation) error
	GetFollowerByID(userID int64) (followerIds []int64, err error)
}
type RelationService struct{}

func NewRelationService() IRelationService {
	return RelationService{}
}

//添加关注
func (f RelationService) AddFollow(rel *model.Relation) (err error) {
	if err = conf.MySQL.Create(&rel).Error; err != nil {
		return err
	}
	return nil
}

//取消关注
func (f RelationService) UnFollow(rel *model.Relation) (err error) {
	if err = conf.MySQL.Where("user_id = ? AND follow_id = ?", rel.UserId, rel.FollowId).
		Delete(&model.Relation{}).Error; err != nil {
		return err
	}
	return nil
}
func (f RelationService) UpdateAdd(rel *model.Relation) error {
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", rel.UserId).
		UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		return err
	}
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", rel.FollowId).
		UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (f RelationService) UpdateSub(rel *model.Relation) error {
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", rel.UserId).
		UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		return err
	}
	if err := conf.MySQL.Model(&model.User{}).
		Where("id = ?", rel.FollowId).
		UpdateColumn("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

//获取关注列表
func (f RelationService) GetFollowByID(userID int64) (followIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Relation{}).
		Select("user_id").
		Where("follow_id = ?", userID).
		Find(&followIds).Error; err != nil {
		return nil, err
	}
	return followIds, err

}

//获取粉丝列表
func (f RelationService) GetFollowerByID(userID int64) (followerIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Relation{}).
		Select("follow_id").
		Where("user_id = ?", userID).
		Find(&followerIds).Error; err != nil {
		return nil, err
	}
	return followerIds, err
}
