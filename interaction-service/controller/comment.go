package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/interaction-service/service"
	"github.com/zsj-dev/DouYin/pb"
)

type CommentServerImpl struct {
	pb.UnimplementedCommentServiceServer
}

var CommentService = service.CommentService{}

func (p *CommentServerImpl) Action(ctx context.Context, req *pb.CommentActionRequest) (*pb.CommentActionResponse, error) {
	if req.ActionType == 1 {
		com := model.Comment{
			UserId:  req.UserID,
			VideoId: req.VideoID,
			Content: req.CommentText,
		}
		if err := CommentService.Create(&com); err != nil {
			return nil, err
		}
		if err := CommentService.UpdateAdd(&com); err != nil {
			return nil, err
		}

		user, err := service.GetUser(req.UserID)
		if err != nil {
			return nil, err
		}
		return &pb.CommentActionResponse{
			Comment: &pb.Comment{
				Id: com.ID,
				User: &pb.User{
					Id:            user.ID,
					Name:          user.Username,
					FollowerCount: user.FollowerCount,
					FollowCount:   user.FollowCount,
					IsFollow:      false,
				},
				Content:    com.Content,
				CreateDate: com.CreatedAt.Format("2006-01-02"),
			},
		}, nil
	} else {

		com := model.Comment{
			UserId:  req.UserID,
			VideoId: req.VideoID,
			BaseModel: model.BaseModel{
				ID: req.CommentID},
		}
		if err := CommentService.Delete(&com); err != nil {
			return nil, err
		}
		if err := CommentService.UpdateSub(&com); err != nil {
			return nil, err
		}
		return &pb.CommentActionResponse{}, nil
	}

}

func (p *CommentServerImpl) List(ctx context.Context, req *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	list, err := CommentService.List(req.VideoId)
	if err != nil {
		return nil, err
	}
	resList := make([]*pb.Comment, 0)

	for _, comment := range list {
		user, err := service.GetUser(comment.UserId)
		if err != nil {
			return nil, err
		}
		isFollow, err := service.IsFollow(comment.UserId, req.UserId)
		if err != nil {
			return nil, err
		}
		resList = append(resList, &pb.Comment{
			Id: comment.ID,
			User: &pb.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowerCount: user.FollowerCount,
				FollowCount:   user.FollowCount,
				IsFollow:      isFollow,
			},
			Content:    comment.Content,
			CreateDate: comment.CreatedAt.Format("2006-01-02"),
		})
	}

	reply := &pb.CommentListResponse{
		List: resList,
	}
	return reply, nil
}
