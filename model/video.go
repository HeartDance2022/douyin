package model

import (
	"douyin/dao"
	"time"
)

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

// CreateVideo Video Post
func CreateVideo(video *Video) error {
	return dao.DB.Create(video).Error
}

// GetPostList GetUploadList
func GetPostList(userid int64) ([]Video, error) {
	var videos []Video
	err := dao.DB.Where("user_id", userid).Find(&videos).Error
	return videos, err
}

// GetVideoById 通过ID获取视频
func GetVideoById(id int64) (*Video, error) {
	var video Video
	err := dao.DB.First(&video, "id = ?", id).Error
	return &video, err
}
