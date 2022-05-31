package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"errors"
	"gorm.io/gorm"
	"time"
)

func Like(like *entity.LikeRequest) (entity.Response, error) {
	userId := like.UserId
	videoId := like.VideoId
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	_, err1 := model.GetVideoById(videoId)
	if err != nil || err1 != nil {
		return util.IDErrorResponse, err
	}

	likeStation, err := model.GetLikedById(userId, videoId)
	//不存在则新建，存在则更新
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//点赞
		if like.ActionType == 1 {
			likeStation.UserID = userId
			likeStation.VideoID = videoId
			likeStation.IsFavorite = true
			likeStation.CreatedAt = time.Now()
			likeStation.UpdatedAt = time.Now()

			err := model.CreateLike(likeStation)
			if err != nil {
				return util.InsertErrorResponse, err
			}
		} else {
			//不存在没办法取消关注
			return util.ServerErrorResponse, err
		}
	} else {
		if like.ActionType == 1 {
			likeStation.IsFavorite = true
			likeStation.UpdatedAt = time.Now()
		} else {
			likeStation.IsFavorite = false
			likeStation.UpdatedAt = time.Now()
		}
		err := model.UpdateLike(likeStation)
		if err != nil {
			return util.UpdateErrorResponse, err
		}
	}

	return util.SuccessResponse, nil
}

// GetFavoriteVideoList 获取userID的点赞视频
func GetFavoriteVideoList(userId int64) (entity.VideoListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	if err != nil {
		return entity.VideoListResponse{Response: util.IDErrorResponse}, err
	}
	videoList, err := model.GetFavoriteVideoList(userId)
	if err != nil {
		return entity.VideoListResponse{Response: util.ServerErrorResponse}, err
	}
	if videoList == nil {
		return entity.VideoListResponse{Response: util.ListNilResponse}, err
	}
	return entity.VideoListResponse{
		Response:  util.SuccessResponse,
		VideoList: videoList,
	}, nil
}
