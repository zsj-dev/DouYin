package model

type Favorite struct {
	BaseModel
	UserId  int64 `gorm:"size:64" json:"user_id"`
	VideoId int64 `gorm:"size:64" json:"video_id"`
}
