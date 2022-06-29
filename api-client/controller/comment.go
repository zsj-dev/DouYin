package controller

import "github.com/gin-gonic/gin"

type ICommentController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type CommentController struct{}

func NewCommentController() ICommentController {
	return CommentController{}
}

func (u CommentController) Action(ctx *gin.Context) {
}

func (u CommentController) List(ctx *gin.Context) {
}
