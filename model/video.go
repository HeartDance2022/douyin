package model

import "time"

type Video struct {
	ID            int64
	UserID        int64
	PlayUrl       string `gorm:"size:1000"`
	CoverUrl      string `gorm:"size:1000"`
	VideoText     string
	FavoriteCount int64
	CommentCount  int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
