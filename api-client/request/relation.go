package request

type RelationActionRequest struct {
	UserId     int64 `json:"user_id" binding:"required"  form:"user_id"`
	ToUserId   int64 `json:"to_user_id"  binding:"required" form:"to_user_id"`
	ActionType int32 `json:"action_type" binding:"required" form:"action_type"`
}
type RelationFollowListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
type RelationFollowerListRequest struct {
	UserId int64 `json:"user_id" binding:"required" form:"user_id"`
}
