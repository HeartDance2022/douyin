package model

import "time"

type Comment struct {
	ID        int64
	UserID    int64
	VideoID   int64
	Content   string
	ParentId  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
