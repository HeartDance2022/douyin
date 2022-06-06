package model

import (
	"douyin/dao"
	"douyin/entity"
	"douyin/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func LikeVideo(userId int64, videoId int64) error {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	//开启redis事务
	redisConn.Send("MULTI")
	videoLikeKey := util.GetVideoLikeKey(videoId)
	userLikedVideoKey := util.GetUserLikedVideoKey(userId)
	_, err := redisConn.Do("SADD", videoLikeKey, userId)
	_, err = redisConn.Do("LPUSH", userLikedVideoKey, videoId)
	// 执行事务
	_, err = redisConn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func UnLikeVideo(userId int64, videoId int64) error {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	//开启redis事务
	redisConn.Send("MULTI")
	videoLikeKey := util.GetVideoLikeKey(videoId)
	userLikedVideoKey := util.GetUserLikedVideoKey(userId)
	_, err := redisConn.Do("SREM", videoLikeKey, userId)
	_, err = redisConn.Do("LREM", userLikedVideoKey, 0, videoId)
	// 执行事务
	_, err = redisConn.Do("EXEC")
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

// GetFavoriteVideoList 获取userID的点赞视频
func GetFavoriteVideoList(userId int64) (videoList []entity.Video, err error) {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err2 := redisConn.Close()
		if err2 != nil {
			panic(err2)
		}
	}(redisConn)
	userLikedVideoKey := util.GetUserLikedVideoKey(userId)
	videoIds, err := redis.Ints(redisConn.Do("LRANGE", userLikedVideoKey, 0, -1))

	for _, v := range videoIds {
		video, err1 := GetVideoById(int64(v))
		if err1 != nil {
			continue
		}
		var isFollowed = HasFollowed(userId, video.UserID)
		//作者信息
		user, _ := GetUserById(video.UserID)
		author := entity.User{
			Id:              user.ID,
			Name:            user.Name,
			FollowCount:     GetFolloweeCount(user.ID),
			FollowerCount:   GetFollowerCount(user.ID),
			IsFollow:        isFollowed,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
		}
		//视频信息
		VideoRep := entity.Video{
			Id:            video.ID,
			Author:        author,
			PlayUrl:       util.ObjGetURL(video.PlayUrl),
			CoverUrl:      util.ObjGetURL(video.CoverUrl),
			FavoriteCount: GetVideoLikedCount(video.ID),
			CommentCount:  video.CommentCount,
			IsFavorite:    true,
			Title:         video.VideoText,
		}
		videoList = append(videoList, VideoRep)
	}
	return videoList, err
}

// FindLikedStatus 点赞状态
func FindLikedStatus(userId int64, videoId int64) bool {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	videoLikeKey := util.GetVideoLikeKey(videoId)
	//是否存在指定key
	flag, err := redis.Int(redisConn.Do("SISMEMBER", videoLikeKey, userId))
	if err != nil {
		return false
	}
	return flag == 1
}

// GetVideoLikedCount 视频被点赞的数量
func GetVideoLikedCount(videoId int64) int64 {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	videoLikeKey := util.GetVideoLikeKey(videoId)
	count, err := redis.Int(redisConn.Do("SCARD", videoLikeKey))
	if err != nil {
		return 0
	}
	return int64(count)
}

// GetUserLikedCount 用户点赞的数量
func GetUserLikedCount(userId int64) int64 {
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	userLikedVideoKey := util.GetUserLikedVideoKey(userId)
	count, err := redis.Int(redisConn.Do("LLEN", userLikedVideoKey))
	if err != nil {
		return 0
	}
	return int64(count)
}

func GetUserTotalLikedCount(userId int64) int64 {
	var sum int64
	redisConn := dao.RedisPool.Get()
	defer func(redisConn redis.Conn) {
		err := redisConn.Close()
		if err != nil {
			panic(err)
		}
	}(redisConn)
	//通过id查找用户所有投稿视频
	videoList, err := GetPostList(userId)
	if err != nil {
		return sum
	}
	for _, video := range videoList {
		videoLikeKey := util.GetVideoLikeKey(video.ID)
		count, _ := redis.Int(redisConn.Do("SCARD", videoLikeKey))
		sum += int64(count)
	}
	return sum
}
