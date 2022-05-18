package service

import (
	"douyin/entity"
	"douyin/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

func Follow(relation *entity.RelationRequest) (entity.Response, error) {
	userId := relation.UserId
	toUserId := relation.ToUserId
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	_, err1 := model.GetUserById(toUserId)
	if err != nil || err1 != nil || userId == toUserId {
		return entity.Response{StatusCode: 400, StatusMsg: "Some users do not exist"}, err
	}
	//拿到关系
	follow, err := model.GetRelationByUserId(userId, toUserId)
	//不存在则新建，存在则更新
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//关注
		if relation.ActionType == 1 {
			follow.IsFollow = true
			follow.FollowerID = userId
			follow.FolloweeID = toUserId
			follow.CreatedAt = time.Now()
			follow.UpdatedAt = time.Now()

			err := model.Create(follow)
			if err != nil {
				return entity.Response{StatusCode: 400, StatusMsg: "Insertion failure"}, err
			}
		} else {
			//不存在关系没办法取消关注
			return entity.Response{StatusCode: 400, StatusMsg: "no relationship"}, err
		}
	} else {
		if relation.ActionType == 1 {
			follow.IsFollow = true
			follow.UpdatedAt = time.Now()
		} else {
			follow.IsFollow = false
			follow.UpdatedAt = time.Now()
		}
		err := model.Update(follow)
		if err != nil {
			return entity.Response{StatusCode: 400, StatusMsg: "update failure"}, err
		}
	}

	return entity.Response{
		StatusCode: 0,
		StatusMsg:  "success",
	}, nil
}

func GetFollowList(userId int64) (entity.FollowListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	if err != nil {
		return entity.FollowListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "userId does not exist"}}, err
	}
	users, err := model.GetFollowList(userId)
	if err != nil || users == nil {
		return entity.FollowListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "get failure"}}, err
	}
	return entity.FollowListResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: users,
	}, nil
}

func GetFollowerList(userId int64) (entity.FollowListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetUserById(userId)
	if err != nil {
		return entity.FollowListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "userId does not exist"}}, err
	}
	users, err := model.GetFollowerList(userId)
	if err != nil || users == nil {
		return entity.FollowListResponse{Response: entity.Response{StatusCode: 400, StatusMsg: "get failure"}}, err
	}
	return entity.FollowListResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
		UserList: users,
	}, nil
}
