package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/controller"
)

func RegisterFeedRouter(r *gin.RouterGroup) {
	feedController := controller.NewFeedController()
	group := r.Group("")
	group.GET("", feedController.Feed)
}
