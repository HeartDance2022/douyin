package model

import (
	"douyin/dao"
	"douyin/entity"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type Follow struct {
	ID         int64
	FolloweeID int64
	FollowerID int64
	IsFollow   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GetRelationByUserId 通过双方ID获取关系
func GetRelationByUserId(userId int64, toUserId int64) (*Follow, error) {
	var relation Follow
	err := dao.DB.Where(&Follow{FollowerID: userId, FolloweeID: toUserId}).First(&relation).Error
	return &relation, err
}

// Create 关注
func Create(follow *Follow) (err error) {
	if err = dao.DB.Create(&follow).Error; err != nil {
		log.Println(err)
	}
	return err
}

func Update(follow *Follow) (err error) {
	return dao.DB.Model(follow).Updates(map[string]interface{}{
		"is_follow": follow.IsFollow,
	}).Error
}

// GetFollowList 获取userID的关注者
func GetFollowList(userId int64) (users []entity.User, err error) {
	var followList []Follow
	err = dao.DB.Where("follower_id = ? AND is_follow = ?", userId, 1).Find(&followList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users, err
	}
	return pakUser(followList, userId, "followee")
}

// GetFollowerList 获取userID的粉丝
func GetFollowerList(userId int64) (users []entity.User, err error) {
	var followerList []Follow
	err = dao.DB.Where("followee_id = ? AND is_follow = ?", userId, 1).Find(&followerList).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return users, err
	}
	return pakUser(followerList, userId, "follower")
}

func pakUser(followerList []Follow, userId int64, name string) (users []entity.User, err error) {
	fmt.Println(followerList, userId)
	for _, followModel := range followerList {
		var userModel *User
		if name == "follower" {
			userModel, err = GetUserById(followModel.FollowerID)
		} else {
			userModel, err = GetUserById(followModel.FolloweeID)
		}
		if err != nil {
			continue
		}
		var user entity.User
		user.Id = userModel.ID
		user.FollowCount = userModel.FollowCount
		user.FollowerCount = userModel.FollowerCount
		follow, err1 := GetRelationByUserId(userId, userModel.ID)
		if err1 != nil {
			user.IsFollow = false
		} else {
			user.IsFollow = follow.IsFollow
		}
		users = append(users, user)
	}
	return users, nil
}

// AfterCreate 更新User follower_count
func (follow *Follow) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	user.ID = follow.FolloweeID

	if err = tx.Model(&user).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		log.Println(err)
	}
	return
}

// AfterUpdate 更新User follower_count
func (follow *Follow) AfterUpdate(tx *gorm.DB) (err error) {
	var user User
	user.ID = follow.FolloweeID
	if err = tx.Model(&user).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", btou(follow.IsFollow))).Error; err != nil {
		log.Println(err)
	}
	return
}

func btou(b bool) int64 {
	if b {
		return 1
	}
	return -1
}
