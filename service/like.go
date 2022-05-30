package service

import (
	"douyin/entity"
	"douyin/model"
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
		return entity.Response{StatusCode: 400, StatusMsg: "video does not exist"}, err
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
				return entity.Response{StatusCode: 400, StatusMsg: "Insertion failure"}, err
			}
		} else {
			//不存在没办法取消关注
			return entity.Response{StatusCode: 400, StatusMsg: "There is no likes"}, err
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
			return entity.Response{StatusCode: 400, StatusMsg: "update failure"}, err
		}
	}

	return entity.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

// GetFavoriteVideoList 获取userID的点赞视频
func GetFavoriteVideoList(userId int64) (entity.VideoListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	if err != nil {
		return entity.VideoListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "userId does not exist"}}, err
	}
	videoList, err := model.GetFavoriteVideoList(userId)
	if err != nil {
		return entity.VideoListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "get failure"}}, err
	}
	if videoList == nil {
		return entity.VideoListResponse{Response: entity.Response{StatusCode: 0, StatusMsg: "粉丝列表为空"}}, err
	}
	return entity.VideoListResponse{
		Response:  entity.Response{StatusCode: 0, StatusMsg: "success"},
		VideoList: videoList,
	}, nil
}
