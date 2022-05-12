package model

import (
	"douyin/dao"
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
