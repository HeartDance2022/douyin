package model

import (
	"douyin/dao"
	"time"
)

type User struct {
	ID              int64
	Password        string
	Name            string
	Avatar          string
	Signature       string
	BackgroundImage string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// GetUserById 通过ID获取用户
func GetUserById(id int64) (*User, error) {
	var user User
	err := dao.DB.First(&user, "id = ?", id).Error
	return &user, err
}

// GetUserByName 通过name获取用户
func GetUserByName(name string) (*User, error) {
	var user User
	err := dao.DB.First(&user, "name = ?", name).Error
	return &user, err
}

// CreateUser 创建新用户
func CreateUser(user *User) error {
	return dao.DB.Create(user).Error
}
