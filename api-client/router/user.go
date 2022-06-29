package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/controller"
	"github.com/zsj-dev/DouYin/api-client/middleware"
)

func RegisterUserRouter(r *gin.RouterGroup) {
	userController := controller.NewUserController()
	r.POST("/login/", userController.Login)
	r.POST("/register/", userController.Register)

	group := r.Group("")
	group.Use(middleware.JWTAuthMiddleware())
	{
		group.GET("", userController.Info)
	}
}
