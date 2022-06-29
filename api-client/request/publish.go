package request

type PublishListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
