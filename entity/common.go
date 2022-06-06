package entity

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	// 下面三项为https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg标注的新增作者信息
	Avatar          string `json:"avatar,omitempty"`
	Signature       string `json:"signature,omitempty"`
	BackgroundImage string `json:"background_image,omitempty"`
	FavoriteCount   int64  `json:"favorite_count,omitempty"`
	TotalFavorited  int64  `json:"total_favorited,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type FeedRequest struct {
	Latest_time time.Time `json:"latest_time"`
}

type RelationRequest struct {
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
	ToUserId   int64  `json:"to_user_id"`
	ActionType int32  `json:"action_type"`
}

type LikeRequest struct {
	UserId     int64  `json:"user_id"`
	Token      string `json:"token"`
	VideoId    int64  `json:"video_id"`
	ActionType int32  `json:"action_type"`
}

type FollowListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
