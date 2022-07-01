package service

import (
	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/interaction-service/conf"
	"github.com/zsj-dev/DouYin/pb"
	"gorm.io/gorm"
)

type FavoriteServerImpl struct {
	pb.UnimplementedFavoriteServiceServer
}

type IFavoriteService interface {
	Delete(fav *model.Favorite) error
	Create(fav *model.Favorite) error
	UpdateLike(fav *model.Favorite) error
	UpdateDisLike(fav *model.Favorite) error
	List(userID int64) (videoIds []int64, err error)
}
type FavoriteService struct{}

func NewFavoriteService() IFavoriteService {
	return FavoriteService{}
}

func (u FavoriteService) Delete(fav *model.Favorite) error {
	if err := conf.MySQL.
		Where("user_id = ? AND video_id = ?", fav.UserId, fav.VideoId).
		Delete(&model.Favorite{}).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) Create(fav *model.Favorite) error {
	if err := conf.MySQL.Create(&fav).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) UpdateLike(fav *model.Favorite) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ?", fav.VideoId).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) UpdateDisLike(fav *model.Favorite) error {
	if err := conf.MySQL.Model(&model.Video{}).
		Where("id = ? AND favorite_count > ?", fav.VideoId, 0).
		UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
		return err
	}
	return nil
}

func (u FavoriteService) List(userID int64) (videoIds []int64, err error) {
	if err := conf.MySQL.Model(&model.Favorite{}).
		Select("video_id").
		Where("user_id = ?", userID).
		Find(&videoIds).Error; err != nil {
		return nil, err
	}

	return videoIds, err
}
