package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/request"
	"github.com/zsj-dev/DouYin/api-client/response"
	"github.com/zsj-dev/DouYin/api-client/service"
	"github.com/zsj-dev/DouYin/pb"
)

type ICommentController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CommentController struct{}

func NewCommentController() ICommentController {
	return CommentController{}
}

func (u CommentController) Action(ctx *gin.Context) {
	payload := request.CommentActionRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	payload.UserId = ctx.GetInt64("user_id")
	if payload.ActionType == 1 {
		resp, err := service.CommentClient.Action(ctx, &pb.CommentActionRequest{
			VideoID:     payload.VideoId,
			UserID:      payload.UserId,
			ActionType:  payload.ActionType,
			CommentText: payload.CommentText,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":       err.Error(),
				"status_code": http.StatusBadRequest,
				"status_msg":  "添加评论操作失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "",
			"comment": response.Comment{
				Id:         resp.Comment.Id,
				Content:    resp.Comment.Content,
				CreateDate: resp.Comment.CreateDate,
				User: response.User{
					Id:            resp.Comment.User.Id,
					Name:          resp.Comment.User.Name,
					FollowCount:   resp.Comment.User.FollowCount,
					FollowerCount: resp.Comment.User.FollowerCount,
					IsFollow:      resp.Comment.User.IsFollow,
				},
			},
		})
	} else {
		_, err := service.CommentClient.Action(ctx, &pb.CommentActionRequest{
			VideoID:    payload.VideoId,
			UserID:     payload.UserId,
			ActionType: payload.ActionType,
			CommentID:  payload.CommentID,
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":       err.Error(),
				"status_code": http.StatusBadRequest,
				"status_msg":  "删除操作失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "",
		})
	}
}

func (u CommentController) List(ctx *gin.Context) {
	payload := request.CommentListRequest{}
	if err := ctx.ShouldBindQuery(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":       err.Error(),
			"status_code": http.StatusBadRequest,
			"status_msg":  "参数错误",
		})
		return
	}
	resp, err := service.CommentClient.List(ctx, &pb.CommentListRequest{
		VideoId: payload.VideoId,
		UserId:  ctx.GetInt64("user_id"),
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"status_msg":  "拉取评论失败",
			"error":       err.Error(),
		})
		return
	}

	// 组装数据返回
	commentList := response.CommentList{}
	for _, comment := range resp.List {
		commentList = append(commentList, response.Comment{
			Id:         comment.Id,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
			User: response.User{
				Id:            comment.User.Id,
				Name:          comment.User.Name,
				FollowCount:   comment.User.FollowCount,
				FollowerCount: comment.User.FollowerCount,
				IsFollow:      comment.User.IsFollow,
			},
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   nil,
		"comment_list": commentList,
	})
}
