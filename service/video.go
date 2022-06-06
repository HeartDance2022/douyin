package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"errors"
	"gorm.io/gorm"
	"log"
	"mime/multipart"
	"strconv"
	"time"
)

func PostList(idStr string, token string) entity.VideoListResponse {
	thisUser := GetLoginUser(token)
	if thisUser == nil {
		return entity.VideoListResponse{Response: util.TokenFailResponse}
	}
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return entity.VideoListResponse{Response: util.ParamErrorResponse}
	}
	user, err := model.GetUserById(userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.VideoListResponse{
				Response: util.IDErrorResponse,
			}
		} else {
			log.Println(err)
			return entity.VideoListResponse{Response: util.ServerErrorResponse}
		}
	}
	var isFollowed = model.HasFollowed(thisUser.ID, user.ID)
	author := entity.User{
		Id:              user.ID,
		Name:            user.Name,
		FollowCount:     model.GetFolloweeCount(user.ID),
		FollowerCount:   model.GetFollowerCount(user.ID),
		IsFollow:        isFollowed,
		Avatar:          user.Avatar,
		FavoriteCount:   model.GetUserLikedCount(user.ID),
		TotalFavorited:  model.GetUserTotalLikedCount(user.ID),
		Signature:       user.Signature,
		BackgroundImage: user.BackgroundImage,
	}
	//通过id查找用户所有投稿视频
	videoList, err := model.GetPostList(user.ID)
	if err != nil {
		return entity.VideoListResponse{}
	}
	resp := entity.VideoListResponse{
		Response:  util.SuccessResponse,
		VideoList: nil,
	}
	for _, video := range videoList {
		resp.VideoList = append(resp.VideoList, mdConv(video, author))
	}
	return resp
}

func mdConv(video model.Video, user entity.User) entity.Video {
	return entity.Video{
		Id:            video.ID,
		Author:        user,
		PlayUrl:       util.ObjGetURL(video.PlayUrl),  //Obtained through pre-signature, private links are not publishe directly
		CoverUrl:      util.ObjGetURL(video.CoverUrl), //Obtained through pre-signature, private links are not publishe directly
		FavoriteCount: model.GetVideoLikedCount(video.ID),
		CommentCount:  video.CommentCount,
	}
}

func PostVideo(form *multipart.Form) entity.Response {
	files := form.File["data"]
	token := form.Value["token"][0]
	title := form.Value["title"][0]
	//Verify Token
	loginUser := GetLoginUser(token)
	if loginUser == nil {
		return util.TokenFailResponse
	}
	for _, file := range files {
		//Playback URL, Screenshot URL
		playUrl, coverUrl, err := util.ObjPost(file)
		if err != nil {
			return util.ServerErrorResponse
		}
		err = model.CreateVideo(&model.Video{
			UserID:    loginUser.ID,
			PlayUrl:   playUrl,
			CoverUrl:  coverUrl,
			VideoText: title,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			return util.ServerErrorResponse
		}
	}
	return util.SuccessResponse
}
