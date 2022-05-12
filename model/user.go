package model

import (
	"douyin/dao"
	"time"
)

type User struct {
	ID            int64
	Password      string
	Name          string
	Description   string
	FollowCount   int64
	FollowerCount int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// GetUserById 通过ID获取用户
func GetUserById(id int64) (*User, error) {
	var user User
	err := dao.DB.First(&user, "id = ?", id).Error
	return &user, err
}
