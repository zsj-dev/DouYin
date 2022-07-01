package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/pb"
	"github.com/zsj-dev/DouYin/relation-service/service"
)

type RelationServiceImpl struct {
	pb.UnimplementedRelationServiceServer
}

var RelationService = service.RelationService{}

func (p *RelationServiceImpl) Action(ctx context.Context, req *pb.RelationActionRequest) (*pb.RelationActionResponse, error) {
	rel := model.Relation{
		UserId:   req.UserID,
		FollowId: req.FollowID,
	}
	if req.ActionType == 1 {
		if err := RelationService.AddFollow(&rel); err != nil {
			return nil, err
		}
		if err := RelationService.UpdateAdd(&rel); err != nil {
			return nil, err
		}

	} else if req.ActionType == 2 {
		if err := RelationService.UnFollow(&rel); err != nil {
			return nil, err
		}
		if err := RelationService.UpdateSub(&rel); err != nil {
			return nil, err
		}
	}
	return &pb.RelationActionResponse{}, nil
}

func (p *RelationServiceImpl) FollowList(ctx context.Context, req *pb.RelationFollowListRequest) (*pb.RelationFollowListResponse, error) {
	followIds, err := RelationService.GetFollowByID(req.UserID)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.User, 0)
	for _, id := range followIds {
		user, err := service.GetUser(id)
		if err != nil {
			return nil, err
		}
		list = append(list, &pb.User{
			Id:            user.ID,
			Name:          user.Username,
			FollowerCount: user.FollowerCount,
			FollowCount:   user.FollowCount,
			IsFollow:      true,
		})

	}
	reply := &pb.RelationFollowListResponse{
		UserList: list,
	}
	return reply, nil
}

func (p *RelationServiceImpl) FollowerList(ctx context.Context, req *pb.RelationFollowerListRequest) (*pb.RelationFollowerListResponse, error) {
	followIds, err := RelationService.GetFollowerByID(req.UserID)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.User, 0)
	for _, id := range followIds {
		user, err := service.GetUser(id)
		if err != nil {
			return nil, err
		}
		isfol, err := service.IsFollow(user.ID, req.UserID)
		if err != nil {
			return nil, err
		}
		list = append(list, &pb.User{
			Id:            user.ID,
			Name:          user.Username,
			FollowerCount: user.FollowerCount,
			FollowCount:   user.FollowCount,
			IsFollow:      isfol,
		})

	}
	reply := &pb.RelationFollowerListResponse{
		UserList: list,
	}
	return reply, nil
}
