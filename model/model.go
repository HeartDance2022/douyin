package model

import "time"

type Video struct {
	Id            int64
	UserId        int64
	PlayUrl       string
	CoverUrl      string
	VideoText     string
	FavoriteCount int64
	CommentCount  int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Comment struct {
	Id        int64
	UserId    int64
	VideoId   int64
	Content   string
	ParentId  int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Id            int64
	Password      string
	Name          string
	Description   string
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Follow struct {
	Id         int64
	FolloweeId int64
	FollowerId int64
	IsFollow   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Like struct {
	Id         int64
	VideoId    int64
	UserId     int64
	IsFavorite bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
