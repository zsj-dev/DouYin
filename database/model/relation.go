package model

type Relation struct {
	BaseModel
	UserId   int64 `gorm:"size:64" json:"user_id"`
	FollowId int64 `gorm:"size:64" json:"follow_id"`
}
