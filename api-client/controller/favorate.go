package controller

import "github.com/gin-gonic/gin"

type IFavoriteController interface {
	Action(ctx *gin.Context)
	List(ctx *gin.Context)
}

type FavoriteController struct{}

func NewFavoriteController() IFavoriteController {
	return FavoriteController{}
}

func (u FavoriteController) Action(ctx *gin.Context) {
}

func (u FavoriteController) List(ctx *gin.Context) {
}
