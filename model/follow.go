package model

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func Following(userId int64, toUserId int64) error {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	//开启redis事务
	redisConn.Send("MULTI")
	followerKey := util.GetFollowerKey(toUserId)
	followeeKey := util.GetFolloweeKey(userId)
	//likedId变成userId的关注者
	_, err := redisConn.Do("SADD", followeeKey, toUserId)
	//userId变成likedId的粉丝
	_, err = redisConn.Do("SADD", followerKey, userId)
	// 执行事务
	_, err = redisConn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func UnFollowing(userId int64, toUserId int64) error {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	//开启redis事务
	redisConn.Send("MULTI")
	followerKey := util.GetFollowerKey(toUserId)
	followeeKey := util.GetFolloweeKey(userId)
	//删除数据
	_, delErr := redisConn.Do("SREM", followeeKey, toUserId)
	_, delErr = redisConn.Do("SREM", followerKey, userId)
	// 执行事务
	_, err := redisConn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return delErr
}

// GetFollowList 获取userID的关注者
func GetFollowList(userId int64, curUserId int64) (users []entity.User, err error) {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err1 := redisConn.Close()
		if err1 != nil {
			panic(err1)
		}
	}(redisConn)
	followeeKey := util.GetFolloweeKey(userId)
	//是否存在指定key
	followeeIds, err := redis.Ints(redisConn.Do("SMEMBERS", followeeKey))
	for _, v := range followeeIds {
		userModel, err1 := GetUserById(int64(v))
		if err1 != nil {
			continue
		}
		var user entity.User
		user.Id = userModel.ID
		user.Name = userModel.Name
		user.FollowCount = GetFolloweeCount(userModel.ID)
		user.FollowerCount = GetFollowerCount(userModel.ID)
		user.IsFollow = HasFollowed(curUserId, int64(v))
		users = append(users, user)
	}
	return users, err
}

// GetFollowerList 获取userID的粉丝
func GetFollowerList(userId int64, curUserId int64) (users []entity.User, err error) {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	followerKey := util.GetFollowerKey(userId)
	//是否存在指定key
	followerIds, err := redis.Ints(redisConn.Do("SMEMBERS", followerKey))
	for _, v := range followerIds {
		userModel, err1 := GetUserById(int64(v))
		if err1 != nil {
			continue
		}
		var user entity.User
		user.Id = userModel.ID
		user.Name = userModel.Name
		user.FollowCount = GetFolloweeCount(userModel.ID)
		user.FollowerCount = GetFollowerCount(userModel.ID)
		user.IsFollow = HasFollowed(curUserId, int64(v))
		users = append(users, user)
	}
	return users, err
}

//HasFollowed 判断当前用户是否已关注该实体
func HasFollowed(userId int64, likedId int64) bool {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	followeeKey := util.GetFolloweeKey(userId)
	//是否存在指定key
	flag, err := redis.Int(redisConn.Do("SISMEMBER", followeeKey, likedId))
	if err != nil {
		return false
	}
	return flag == 1
}

// GetFolloweeCount 用户关注的数量
func GetFolloweeCount(userId int64) int64 {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	followeeKey := util.GetFolloweeKey(userId)
	count, err := redis.Int(redisConn.Do("SCARD", followeeKey))
	if err != nil {
		return 0
	}
	return int64(count)
}

// GetFollowerCount 用户粉丝数量
func GetFollowerCount(userId int64) int64 {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	followerKey := util.GetFollowerKey(userId)
	count, err := redis.Int(redisConn.Do("SCARD", followerKey))
	if err != nil {
		return 0
	}
	return int64(count)
}
