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
		isFavorite = model.FindLikedStatus(user.ID, i.Video.ID)
		isFollow = model.HasFollowed(user.ID, i.User.ID)
	}

	return entity.Video{
		Id: i.Video.ID,
		Author: entity.User{
			Id:              i.User.ID,
			Name:            i.User.Name,
			FollowCount:     model.GetFolloweeCount(i.User.ID),
			FollowerCount:   model.GetFollowerCount(i.User.ID),
			TotalFavorited:  model.GetUserTotalLikedCount(i.User.ID),
			Avatar:          i.User.Avatar,
			FavoriteCount:   model.GetUserLikedCount(i.User.ID),
			IsFollow:        isFollow,
			Signature:       i.User.Signature,
			BackgroundImage: i.User.BackgroundImage,
		},
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		PlayUrl: util.ObjGetURL(i.Video.PlayUrl),
		//通过预签名方式访问私有读写存储桶,不直接存储永久公有url
		CoverUrl:      util.ObjGetURL(i.Video.CoverUrl),
		FavoriteCount: model.GetVideoLikedCount(i.Video.ID),
		CommentCount:  i.Video.CommentCount,
		Title:         i.Video.VideoText,
		IsFavorite:    isFavorite,
	}
}
