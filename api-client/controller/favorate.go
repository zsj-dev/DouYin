package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/request"
	"github.com/zsj-dev/DouYin/api-client/response"
	"github.com/zsj-dev/DouYin/api-client/service"
	"github.com/zsj-dev/DouYin/pb"
)

type IFavoriteController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type FavoriteController struct{}

func NewFavoriteController() IFavoriteController {
	return FavoriteController{}
}

func (u FavoriteController) Action(ctx *gin.Context) {
	payload := request.FavoriteActionRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	payload.UserId = ctx.GetInt64("user_id")
	_, err := service.FavoriteClient.Action(ctx, &pb.FavoriteActionRequest{
		UserID:     payload.UserId,
		VideoID:    payload.VideoId,
		ActionType: payload.ActionType,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "点赞操作失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "",
	})
}

func (u FavoriteController) List(ctx *gin.Context) {
	payload := request.FavoriteListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.FavoriteClient.List(ctx, &pb.FavoriteListRequest{
		UserID: payload.UserId,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取操作失败",
			"error":       err.Error(),
		})
		return
	}

	if resp.List == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  nil,
			"video_list":  nil,
		})
		return
	}

	// 组装数据返回
	videoList := response.VideoList{}
	for _, video := range resp.List {
		videoList = append(videoList, response.Video{
			Author: response.User{
				Id:            video.Author.Id,
				Name:          video.Author.Name,
				FollowCount:   video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow:      video.Author.IsFollow,
			},
			Id:            video.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			Title:         video.Title,
			IsFavorite:    video.IsFavorate,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  nil,
		"video_list":  videoList,
	})
}
