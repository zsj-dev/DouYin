package request

type CommentActionRequest struct {
	UserId      int64  `json:"user_id" form:"user_id"`
	VideoId     int64  `json:"video_id" binding:"required" form:"video_id"`
	ActionType  int64  `json:"action_type" binding:"required" form:"action_type"`
	CommentText string `json:"comment_text,omitempty" form:"comment_text"`
	CommentID   int64  `json:"comment_id,omitempty"  form:"comment_id"`
}

type CommentListRequest struct {
	VideoId int64 `json:"video_id" binding:"required"  form:"video_id"`
}
