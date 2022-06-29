package request

type FeedRequest struct {
	LatestTime int64 `json:"latest_time,omitempty"  form:"latest_time"`
}
