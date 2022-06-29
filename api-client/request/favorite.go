package request

type FavoriteActionRequest struct {
	UserId     int64 `json:"user_id" form:"user_id"`
	VideoId    int64 `json:"video_id" binding:"required" form:"video_id"`
	ActionType int64 `json:"action_type" binding:"required" form:"action_type"`
}

type FavoriteListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
