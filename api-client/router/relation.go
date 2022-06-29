package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/controller"
	"github.com/zsj-dev/DouYin/api-client/middleware"
)

func RegisterRelationRouter(r *gin.RouterGroup) {
	relationController := controller.NewRelationController()

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/action/", relationController.Action)
		group.GET("/follow/list/", relationController.FollowList)
		group.GET("/follower/list/", relationController.FollowerList)
	}
}
