package controller

import (
	"context"

	"github.com/zsj-dev/DouYin/basic-service/service"
	"github.com/zsj-dev/DouYin/database/model"
	"github.com/zsj-dev/DouYin/pb"
	"golang.org/x/crypto/bcrypt"
)

type UserServerImpl struct {
	pb.UnimplementedUserServiceServer
}

var UserService = service.UserService{}

func (p *UserServerImpl) Get(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetResponse, error) {
	isFollow, err := service.IsFollow(req.UserId, req.SeeId)
	if err != nil {
		return nil, err
	}
	user, err := service.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}
	var reply pb.UserGetResponse
	if user != nil {
		reply = pb.UserGetResponse{
			User: &pb.User{
				Id:            user.ID,
				Name:          user.Username,
				FollowCount:   user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow:      isFollow,
			},
		}
	} else {
		reply = pb.UserGetResponse{
			User: &pb.User{
				Id:            0,
				Name:          "",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			},
		}
	}

	return &reply, nil
}

func (p *UserServerImpl) Register(ctx context.Context, req *pb.UserRegisterRequest) (*pb.UserRegisterResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := UserService.Create(&user); err != nil {
		return nil, err
	}
	reply := &pb.UserRegisterResponse{
		UserID: user.ID,
	}
	return reply, nil
}

func (p *UserServerImpl) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	user := model.User{}
	if err := UserService.FindByUsername(&user, req.Username); err != nil {
		return nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	reply := &pb.UserLoginResponse{
		UserID: user.ID,
	}
	return reply, nil
}
