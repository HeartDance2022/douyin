package model

import "time"

type Like struct {
	ID         int64
	VideoID    int64
	UserID     int64
	IsFavorite bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
