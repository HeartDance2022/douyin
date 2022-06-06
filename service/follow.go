package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
)

func Follow(relation *entity.RelationRequest) (entity.Response, error) {
	userId := relation.UserId
	toUserId := relation.ToUserId
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	_, err1 := model.GetUserById(toUserId)
	if err != nil || err1 != nil || userId == toUserId {
		return util.TokenFailResponse, err
	}
	//是否关注
	hasFollowed := model.HasFollowed(userId, toUserId)
	if hasFollowed {
		if relation.ActionType == 1 {
			//点过关注还能关注
			return util.ServerErrorResponse, err
		} else {
			delErr := model.UnFollowing(userId, toUserId)
			if delErr != nil {
				return util.ServerErrorResponse, err
			}
		}
	} else {
		if relation.ActionType == 1 {
			err = model.Following(userId, toUserId)
			if err != nil {
				return util.ServerErrorResponse, err
			}
		} else {
			//没点关注还能取消关注
			return util.ServerErrorResponse, err
		}
	}

	return util.SuccessResponse, nil
}

func GetFollowList(userId int64, curUserId int64) (entity.FollowListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	_, err1 := model.GetUserById(curUserId)
	if err != nil || err1 != nil {
		return entity.FollowListResponse{Response: util.TokenFailResponse}, err
	}
	users, err := model.GetFollowList(userId, curUserId)
	if err != nil {
		return entity.FollowListResponse{Response: util.ServerErrorResponse}, err
	}
	return entity.FollowListResponse{
		Response: util.SuccessResponse,
		UserList: users,
	}, nil
}

func GetFollowerList(userId int64, curUserId int64) (entity.FollowListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	_, err1 := model.GetUserById(curUserId)
	if err != nil || err1 != nil {
		return entity.FollowListResponse{Response: util.TokenFailResponse}, err
	}
	users, err := model.GetFollowerList(userId, curUserId)
	if err != nil {
		return entity.FollowListResponse{Response: util.ServerErrorResponse}, err
	}
	return entity.FollowListResponse{
		Response: util.SuccessResponse,
		UserList: users,
	}, nil
}
