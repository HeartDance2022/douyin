package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"time"
)

// usersLoginInfo 使用map存储登录的用户信息，key值为util.Hash(username+password)
// 注意：map在每次运行时会被清空
var usersLoginInfo = map[string]*model.User{}

func Login(username string, password string) entity.UserLoginResponse {
	// 根据name找到user
	user, err := model.GetUserByName(username)
	// 如果没找到或者其他错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没找到
			return entity.UserLoginResponse{
				Response: entity.Response{
					StatusCode: 400,
					StatusMsg:  "用户名或密码错误",
				},
			}
		} else {
			// 其他错误
			log.Println(err)
			return entity.UserLoginResponse{Response: util.ServerErrorResponse}
		}
	}
	// 检查password
	if user.Password != password {
		return entity.UserLoginResponse{
			Response: entity.Response{
				StatusCode: 400,
				StatusMsg:  "用户名或密码错误",
			},
		}
	}
	// 登录成功
	hash := util.Hash(username + password)
	usersLoginInfo[hash] = user
	return entity.UserLoginResponse{
		Response: util.SuccessResponse,
		UserId:   user.ID,
		Token:    hash,
	}
}

// GetLoginUser 通过token返回用户信息
func GetLoginUser(token string) *model.User {
	return usersLoginInfo[token]
}

func UserInfo(idStr string, token string) entity.UserResponse {
	// 检查token有效性
	thisUser := GetLoginUser(token)
	if thisUser == nil {
		return entity.UserResponse{
			Response: util.TokenFailResponse,
			User:     entity.User{},
		}
	}
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return entity.UserResponse{Response: util.ParamErrorResponse}
	}
	// 通过id查找user
	user, err := model.GetUserById(userId)
	// 如果没找到或者其他错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 没找到
			return entity.UserResponse{
				Response: entity.Response{
					StatusCode: 400,
					StatusMsg:  "userId无效",
				},
			}
		} else {
			// 其他错误
			log.Println(err)
			return entity.UserResponse{Response: util.ServerErrorResponse}
		}
	}
	// 查找关注信息
	relation, err := model.GetRelationByUserId(thisUser.ID, user.ID)
	var isFollowed = true
	if err != nil || !relation.IsFollow {
		isFollowed = false
	}
	// 成功
	return entity.UserResponse{
		Response: util.SuccessResponse,
		User: entity.User{
			Id:            user.ID,
			Name:          user.Name,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollowed,
		},
	}
}

func Register(username string, password string) entity.UserLoginResponse {
	// 先判断用户名或密码是否为空
	if username == "" || password == "" {
		return entity.UserLoginResponse{Response: entity.Response{
			StatusCode: 400,
			StatusMsg:  "用户名和密码不能为空",
		}}
	}
	// 创建user对象
	user := &model.User{
		Password:  password,
		Name:      username,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := model.CreateUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return entity.UserLoginResponse{Response: entity.Response{
				StatusCode: 400,
				StatusMsg:  "用户名不能重复",
			}}
		}
		fmt.Printf("%v\n", err)
		return entity.UserLoginResponse{Response: util.ServerErrorResponse}
	}
	// 取出创建好的user
	user, err = model.GetUserByName(user.Name)
	if err != nil {
		fmt.Printf("%v\n", err)
		return entity.UserLoginResponse{Response: util.ServerErrorResponse}
	}
	// 注册成功
	hash := util.Hash(username + password)
	usersLoginInfo[hash] = user
	return entity.UserLoginResponse{
		Response: util.SuccessResponse,
		UserId:   user.ID,
		Token:    hash,
	}
}
