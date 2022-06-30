package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/basic-service/service"
	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/pb"
)

type PublishServerImpl struct {
	pb.UnimplementedPublishServiceServer
}

var PublishService = service.PublishService{}

func (p *PublishServerImpl) Action(ctx context.Context, req *pb.PublishActionRequest) (*pb.PublishActionResponse, error) {
	m := model.Video{
		AuthorID:      req.AuthorID,
		Title:         req.Title,
		PlayURL:       req.PlayUrl,
		CoverURL:      req.CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
	}
	_, err := PublishService.Create(&m)
	if err != nil {
		return nil, err
	}

	reply := &pb.PublishActionResponse{}
	return reply, nil
}

func (p *PublishServerImpl) List(ctx context.Context, req *pb.PublishListRequest) (*pb.PublishListResponse, error) {
	resp, err := PublishService.List(req.UserId)
	if err != nil {
		return nil, err
	}

	list := make([]*pb.Video, 0)
	for _, video := range resp {
		user, err := service.GetUser(video.AuthorID)
		if err != nil {
			return nil, err
		}
		isfollow, err := service.IsFollow(user.ID, req.UserId)
		if err != nil {
			return nil, err
		}
		isfav, err := service.IsFav(req.UserId, video.ID)
		if err != nil {
			return nil, err
		}
		list = append(list, &pb.Video{
			Id: video.ID,
			Author: &pb.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
				IsFollow:      isfollow,
			},
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorate:    isfav,
		})
	}
	reply := &pb.PublishListResponse{
		List: list,
	}
	return reply, nil
}
