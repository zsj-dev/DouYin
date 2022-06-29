package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/controller"
	"github.com/zsj-dev/DouYin/api-client/middleware"
)

func RegisterFavoriteRouter(r *gin.RouterGroup) {
	favController := controller.NewFavoriteController()
	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.POST("/action/", favController.Action)
		group.GET("/list/", favController.List)
	}
}
