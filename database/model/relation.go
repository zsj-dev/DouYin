package model

type Relation struct {
	BaseModel
	//被关注人
	UserId int64 `gorm:"size:64" json:"user_id"`
	//关注人
	FollowId int64 `gorm:"size:64" json:"follow_id"`
}
