package initialization

import "github.com/zsj-dev/DouYin/api-client/service"

func SetupService() {
	service.UserConn()
	service.FavoriteConn()
	service.CommentConn()
	service.RelationConn()
	service.PublishConn()
	service.FeedConn()
}
