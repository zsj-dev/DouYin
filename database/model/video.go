package model

type Video struct {
	BaseModel
	AuthorID      int64  `gorm:"column:author_id;not null" json:"author_id"`
	PlayURL       string `gorm:"column:play_url;not null"  json:"play_url"`
	CoverURL      string `gorm:"column:cover_url;not null" json:"cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" json:"comment_count"`
	Title         string `gorm:"column:title;not null" json:"title"`
}
