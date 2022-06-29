package controller

import "github.com/gin-gonic/gin"

type IPublishController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type PublishController struct{}

func NewPublishController() IPublishController {
	return PublishController{}
}

func (u PublishController) Action(ctx *gin.Context) {

}

func (u PublishController) List(ctx *gin.Context) {
}
