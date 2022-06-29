package model

type User struct {
	BaseModel
	Username      string `gorm:"size:64;not null;uniqueIndex" json:"username"`
	Password      string `gorm:"size:255;not null" json:"-"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}
