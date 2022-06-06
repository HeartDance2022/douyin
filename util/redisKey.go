package util

import "strconv"

const SPLIT string = ":"

// UserLikedVideo 用户点赞的视频
const UserLikedVideo string = "user:like:video"

//VideoLike 视频的点赞
const VideoLike string = "like:video"

// FOLLOWER 粉丝
const FOLLOWER string = "follower"

// FOLLOWEE 关注目标
const FOLLOWEE string = "followee"

// GetUserLikedVideoKey 某个用户被喜欢的视频
func GetUserLikedVideoKey(userId int64) string {
	return UserLikedVideo + SPLIT + strconv.FormatInt(userId, 10)
}

// GetVideoLikeKey 某个用户被喜欢的视频
func GetVideoLikeKey(videoId int64) string {
	return VideoLike + SPLIT + strconv.FormatInt(videoId, 10)
}

// GetFolloweeKey 获取用户关注的人
func GetFolloweeKey(userId int64) string {
	return FOLLOWEE + SPLIT + strconv.FormatInt(userId, 10)
}

// GetFollowerKey 获取用户粉丝
func GetFollowerKey(userId int64) string {
	return FOLLOWER + SPLIT + strconv.FormatInt(userId, 10)
}
