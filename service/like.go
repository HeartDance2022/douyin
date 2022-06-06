package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
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
	hasLike := model.FindLikedStatus(userId, videoId)
	//不存在则新建，存在则更新
	if hasLike {
		//点赞
		if like.ActionType == 1 {
			//点过赞还能点赞
			return util.InsertErrorResponse, err
		} else {
			delErr := model.UnLikeVideo(userId, videoId)
			if delErr != nil {
				return util.ServerErrorResponse, err
			}
		}
	} else {
		if like.ActionType == 1 {
			err = model.LikeVideo(userId, videoId)
			if err != nil {
				return util.ServerErrorResponse, err
			}
		} else {
			return util.ServerErrorResponse, err
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
