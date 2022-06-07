package model

import (
	"douyin/dao"
	"douyin/entity"
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Comment struct {
	ID        int64
	UserID    int64
	VideoID   int64
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// getCommentById 通过id查找
func GetCommentById(commentId int64) (*Comment, error) {
	var comment Comment
	err := dao.DB.First(&comment, "id =? ", commentId).Error
	return &comment, err
}

//发布评论
func AddComment(comment *Comment) (err error) {
	if err = dao.DB.Create(&comment).Error; err != nil {
		log.Println(err)
	}
	return err
}

func FindCommentsByVideoId(videoId int64) (con []Comment, err error) {
	var commentList []Comment
	err = dao.DB.Where("video_id = ?", videoId).Find(&commentList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return commentList, err
	}
	return commentList, err
}

//根据id删除
func DeleteCommentByCommentId(commentId int64) (err error) {
	comment, _ := GetCommentById(commentId)
	if err = dao.DB.Delete(&comment).Error; err != nil {
		log.Println(err)
	}
	return err
}

//评论列表
func GetCommentList(videoId int64, userId int64) (CommentList []entity.Comment, err error) {
	s := make([]int64, 100)
	//获取到列表
	comments, err := FindCommentsByVideoId(videoId)
	log.Println("comments", comments)
	for _, comment := range comments {
		//将列表里的所有commentId放到切片里
		s = append(s, comment.ID)
	}
	for _, v := range s {
		if v != 0 {
			comment, err1 := GetCommentById(v)
			if err1 != nil {
				continue
			}
			//第一个是登录的userId从token中取,第二个是评论人的userId
			var isFollowed = HasFollowed(userId, comment.UserID)
			//评论人信息
			user1, _ := GetUserById(comment.UserID)
			user2 := entity.User{
				Id:            user1.ID,
				Name:          user1.Name,
				FollowCount:   GetFolloweeCount(user1.ID),
				FollowerCount: GetFollowerCount(user1.ID),
				IsFollow:      isFollowed,
			}
			//评论
			CommentRep := entity.Comment{
				Id:         comment.ID,
				User:       user2,
				Content:    comment.Content,
				CreateDate: comment.CreatedAt.Format("2006-01-02 15:04:05"),
			}
			CommentList = append(CommentList, CommentRep)
		}
	}
	return CommentList, err
}

func (comment *Comment) AfterCreate(tx *gorm.DB) (err error) {
	var videos Video
	videos.ID = comment.VideoID
	if err = tx.Model(&videos).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		log.Println(err)
	}
	return
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var videos Video
	videos.ID = comment.VideoID
	if err = tx.Model(&videos).UpdateColumn("comment_count", gorm.Expr("comment_count + ?", -1)).Error; err != nil {
		log.Println(err)
	}
	return
}
