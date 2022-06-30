package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/basic-service/service"
	"github.com/zsj-dev/DouYin/pb"
)

type FeedServiceImpl struct {
	pb.UnimplementedFeedServiceServer
}

var FeedService = service.FeedService{}

func (p *FeedServiceImpl) Feed(ctx context.Context, req *pb.FeedRequest) (*pb.FeedResponse, error) {
	resp, err := FeedService.GetVideos(req.LatestTime)
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

	reply := &pb.FeedResponse{
		List:     list,
		NextTime: resp[0].CreatedAt.Unix(),
	}
	return reply, nil
}
