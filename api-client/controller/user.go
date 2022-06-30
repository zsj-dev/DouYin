package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/request"
	"github.com/zsj-dev/DouYin/api-client/response"
	"github.com/zsj-dev/DouYin/api-client/service"
	"github.com/zsj-dev/DouYin/api-client/util"
	"github.com/zsj-dev/DouYin/pb"
	"gorm.io/gorm"
)

type IUserController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	Info(ctx *gin.Context)
}
type UserController struct{}

func NewUserController() IUserController {
	return UserController{}
}

func (u UserController) Login(ctx *gin.Context) {
	payload := request.UserLoginRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.UserClient.Login(ctx, &pb.UserLoginRequest{
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error":       err.Error(),
				"status_code": http.StatusNotFound,
				"status_msg":  "未找到记录",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "用户名或密码错误",
		})
		return
	}
	token, err := util.GenerateToken(resp.UserID, payload.Username, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusInternalServerError,
			"status_msg":  "token 生成失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"user_id":     resp.UserID,
		"token":       token,
	})
}

func (u UserController) Register(ctx *gin.Context) {
	payload := request.UserRegisterRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.UserClient.Register(ctx, &pb.UserRegisterRequest{
		Username: payload.Username,
		Password: payload.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "创建用户失败",
			"error":       err.Error(),
		})
		return
	}
	token, err := util.GenerateToken(resp.UserID, payload.Username, payload.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"user_id":     resp.UserID,
		"token":       token,
	})
}

func (u UserController) Info(ctx *gin.Context) {

	payload := request.UserGetRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.UserClient.Get(ctx, &pb.UserGetRequest{
		UserId: ctx.GetInt64("user_id"),
		SeeId:  payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "查询用户信息失败",
			"error":       err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"user": response.User{
			Id:            resp.User.Id,
			Name:          resp.User.Name,
			FollowCount:   resp.User.FollowCount,
			FollowerCount: resp.User.FollowerCount,
			IsFollow:      resp.User.IsFollow,
		},
	})
}
