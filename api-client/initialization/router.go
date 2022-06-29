package initialization

import (
	"github.com/gin-gonic/gin"
	"github.com/zsj-dev/DouYin/api-client/router"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/douyin")
	router.RegisterUserRouter(apiGroup.Group("/user"))
	router.RegisterFavoriteRouter(apiGroup.Group("/favorite"))
	router.RegisterCommentRouter(apiGroup.Group("/comment"))
	router.RegisterRelationRouter(apiGroup.Group("/relation"))
	router.RegisterPublishRouter(apiGroup.Group("/publish"))
	router.RegisterFeedRouter(apiGroup.Group("/feed"))
	return r
}
