package service

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"time"
)

var videoLimit = 30

type info struct {
	model.Video
	model.User
}

func Feed(latestTime time.Time) (entity.FeedResponse, error) {
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
		videos = append(videos, modelToc(i))
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
		NextTime:  time.Now().Unix(),
	}, nil
}

func modelToc(i info) entity.Video {
	return entity.Video{
		Id: i.Video.ID,
		Author: entity.User{
			Id:            i.User.ID,
			Name:          i.User.Name,
			FollowCount:   i.User.FollowCount,
			FollowerCount: i.User.FollowerCount,
		},
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		PlayUrl: util.ObjGetURL(i.Video.PlayUrl),
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		CoverUrl:      util.ObjGetURL(i.Video.CoverUrl),
		FavoriteCount: i.Video.FavoriteCount,
		CommentCount:  i.Video.CommentCount,
	}
}
