package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/interaction-service/service"
	"github.com/zsj-dev/DouYin/pb"
)

type FavoriteServerImpl struct {
	pb.UnimplementedFavoriteServiceServer
}

var FavoriteService = service.FavoriteService{}

func (p *FavoriteServerImpl) Action(ctx context.Context, req *pb.FavoriteActionRequest) (*pb.FavoriteActionResponse, error) {
	fav := model.Favorite{
		UserId:  req.UserID,
		VideoId: req.VideoID,
	}
	// 点赞
	if req.ActionType == 1 {

		if err := FavoriteService.Create(&fav); err != nil {
			return nil, err
		}
		if err := FavoriteService.UpdateLike(&fav); err != nil {
			return nil, err
		}

		// 取消点赞
	} else if req.ActionType == 2 {
		if err := FavoriteService.Delete(&fav); err != nil {
			return nil, err
		}
		if err := FavoriteService.UpdateDisLike(&fav); err != nil {
			return nil, err
		}
	}

	return &pb.FavoriteActionResponse{}, nil
}

func (p *FavoriteServerImpl) List(ctx context.Context, req *pb.FavoriteListRequest) (*pb.FavoriteListResponse, error) {
	videoids, err := FavoriteService.List(req.UserID)
	if err != nil {
		return nil, err
	}
	list, err := service.GetVideo(videoids)
	if err != nil {
		return nil, err
	}
	resList := make([]*pb.Video, 0)

	for _, video := range list {
		user, err := service.GetUser(video.AuthorID)
		if err != nil {
			return nil, err
		}
		isFollow, err := service.IsFollow(user.ID, req.UserID)
		if err != nil {
			return nil, err
		}

		resList = append(resList, &pb.Video{
			Id: video.ID,
			Author: &pb.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
				IsFollow:      isFollow,
			},
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorate:    true,
		})

	}

	reply := &pb.FavoriteListResponse{
		List: resList,
	}
	return reply, nil
}
