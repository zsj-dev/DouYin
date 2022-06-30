package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/request"
	"github.com/zsj-dev/DouYin/api-client/response"
	"github.com/zsj-dev/DouYin/api-client/service"
	"github.com/zsj-dev/DouYin/pb"
)

type IRelationController interface {
	Action(ctx *gin.Context)
	FollowList(ctx *gin.Context)
	FollowerList(ctx *gin.Context)
}
type RelationController struct{}

func NewRelationController() IRelationController {
	return RelationController{}
}

func (u RelationController) Action(ctx *gin.Context) {
	payload := request.RelationActionRequest{}
	payload.UserId = ctx.GetInt64("user_id")
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	_, err := service.RelationClient.Action(ctx, &pb.RelationActionRequest{
		UserID:     payload.UserId,
		FollowID:   payload.FollowId,
		ActionType: payload.ActionType,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "关注失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
	})

}

func (u RelationController) FollowList(ctx *gin.Context) {
	payload := request.RelationFollowListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.RelationClient.FollowList(ctx, &pb.RelationFollowListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取关注列表失败",
		})
	}

	followList := response.UserList{}
	for _, user := range resp.UserList {

		followList = append(followList, response.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_list":   followList,
	})
}

func (u RelationController) FollowerList(ctx *gin.Context) {
	payload := request.RelationFollowerListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.RelationClient.FollowerList(ctx, &pb.RelationFollowerListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取粉丝列表失败",
		})
	}

	followList := response.UserList{}
	for _, user := range resp.UserList {

		followList = append(followList, response.User{
			Id:            user.Id,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      user.IsFollow,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
		"user_list":   followList,
	})
}
