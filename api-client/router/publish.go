package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/controller"
	"github.com/zsj-dev/DouYin/api-client/middleware"
)

func RegisterPublishRouter(r *gin.RouterGroup) {
	publishController := controller.NewPublishController()
	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("action/", publishController.Action)
		group.GET("list/", publishController.List)
	}
}
