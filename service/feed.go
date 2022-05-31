package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"time"
)

var videoLimit = 10

type info struct {
	model.Video
	model.User
}

func Feed(latestTime time.Time, user *model.User) (entity.FeedResponse, error) {
	rows, err := dao.DB.Debug().Table("videos").
		Where("videos.updated_at < ?", latestTime).
		Select("videos.*, users.*").
		Joins("left join users on users.id = videos.user_id").
		Limit(videoLimit).
		Rows()
	if err != nil {
		return entity.FeedResponse{Response: entity.Response{StatusCode: 400}}, err
	}

	var videos []entity.Video
	for rows.Next() {
		var i info
		dao.DB.ScanRows(rows, &i)
		videos = append(videos, modelToc(i, user))
	}

	if len(videos) < 1 {
		return entity.FeedResponse{Response: entity.Response{
			StatusCode: 400,
			StatusMsg:  "No video yet"}}, nil
	}
	return entity.FeedResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		VideoList: videos,
		NextTime:  time.Now().UnixNano() / 1e6,
	}, nil
}

func modelToc(i info, user *model.User) entity.Video {
	isFavorite := false
	isFollow := false
	if user != nil {
		like, err := model.GetLikedById(user.ID, i.Video.ID)
		if err == nil && like.IsFavorite {
			isFavorite = true
		}

		follow, err := model.GetRelationByUserId(user.ID, i.User.ID)
		if err == nil && follow.IsFollow {
			isFollow = true
		}
	}

	return entity.Video{
		Id: i.Video.ID,
		Author: entity.User{
			Id:            i.User.ID,
			Name:          i.User.Name,
			FollowCount:   i.User.FollowCount,
			FollowerCount: i.User.FollowerCount,
			IsFollow:      isFollow,
		},
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		PlayUrl: util.ObjGetURL(i.Video.PlayUrl),
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		CoverUrl:      util.ObjGetURL(i.Video.CoverUrl),
		FavoriteCount: i.Video.FavoriteCount,
		CommentCount:  i.Video.CommentCount,
		Title:         i.Video.VideoText,
		IsFavorite:    isFavorite,
	}
}
