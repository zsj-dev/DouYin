package controller

import "github.com/gin-gonic/gin"

type IFeedController interface {
	Feed(ctx *gin.Context)
}
type FeedController struct {
}

func NewFeedController() IFeedController {
	return FeedController{}
}
func (u FeedController) Feed(ctx *gin.Context) {
}
