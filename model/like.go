package model

import (
	"douyin/dao"
	"douyin/entity"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Like struct {
	ID         int64
	VideoID    int64
	UserID     int64
	IsFavorite bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// GetLikedById 通过ID判断点赞状态
func GetLikedById(userId int64, videoId int64) (*Like, error) {
	var like Like
	err := dao.DB.Where(&Like{UserID: userId, VideoID: videoId}).First(&like).Error
	return &like, err
}

// CreateLike 点赞
func CreateLike(like *Like) (err error) {
	if err = dao.DB.Create(&like).Error; err != nil {
		log.Println(err)
	}
	return err
}

func UpdateLike(like *Like) (err error) {
	return dao.DB.Model(like).Updates(map[string]interface{}{
		"is_favorite": like.IsFavorite,
	}).Error
}

// GetFavoriteVideoList 获取userID的点赞视频
func GetFavoriteVideoList(userId int64) (videoList []entity.Video, err error) {
	var likeList []Like
	err = dao.DB.Where("user_id = ? AND is_favorite = ?", userId, 1).Find(&likeList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return videoList, err
	}
	for _, likeModel := range likeList {
		video, err := GetVideoById(likeModel.VideoID)
		if err != nil {
			continue
		}
		relation, err := GetRelationByUserId(userId, video.UserID)
		var isFollowed = true
		if err != nil || !relation.IsFollow {
			isFollowed = false
		}
		//作者信息
		user, err := GetUserById(video.UserID)
		author := entity.User{
			Id:              user.ID,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        isFollowed,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
		}
		//视频信息
		VideoRep := entity.Video{
			Id:            video.ID,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.VideoText,
		}
		videoList = append(videoList, VideoRep)
	}
	return videoList, nil
}

// AfterCreate 更新Video favorite_count
func (like *Like) AfterCreate(tx *gorm.DB) (err error) {
	video := Video{ID: like.VideoID}
	if err = tx.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
		log.Println(err)
	}
	return
}

// AfterUpdate 更新Video favorite_count
func (like *Like) AfterUpdate(tx *gorm.DB) (err error) {
	video := Video{ID: like.VideoID}
	if err = tx.Model(&video).UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", btou(like.IsFavorite))).Error; err != nil {
		log.Println(err)
	}
	return
}
