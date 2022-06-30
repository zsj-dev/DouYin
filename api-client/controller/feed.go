package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/response"
	"github.com/zsj-dev/DouYin/api-client/service"
	"github.com/zsj-dev/DouYin/pb"
)

type IFeedController interface {
	Feed(ctx *gin.Context)
}
type FeedController struct {
}

func NewFeedController() IFeedController {
	return FeedController{}
}
func (u FeedController) Feed(ctx *gin.Context) {
	var latestTime, nextTime int64
	latestTime_ := ctx.Query("latest_time")
	if latestTime_ != "" {
		latestTime, _ = strconv.ParseInt(latestTime_, 10, 64)
	} else {
		latestTime = 0
	}
	resp, err := service.FeedClient.Feed(ctx, &pb.FeedRequest{
		LatestTime: latestTime,
		UserId:     ctx.GetInt64("user_id"),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取操作失败",
		})
	}
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
		"status_msg":  "",
		"next_time":   nextTime,
		"video_list":  videoList,
	})
}
