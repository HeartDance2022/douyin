package service

import (
	"douyin/entity"
	"douyin/model"
	"douyin/util"
	"time"
)

func Comment(commentRequest entity.CommentRequest, userId int64) (entity.CommentResponse, error) {
	actionType := commentRequest.ActionType
	videoId := commentRequest.VideoId
	ContentText := commentRequest.CommentText
	if actionType == 1 {

		_, err := model.GetVideoById(videoId)
		_, err1 := model.GetUserById(userId)
		if err != nil || err1 != nil {
			return entity.CommentResponse{
				Response: entity.Response{StatusCode: 400, StatusMsg: "video is not exist or user is not exist"},
				Comment:  entity.Comment{},
			}, err
		}
		// 创建Comment对象
		Comment := &model.Comment{
			UserID:    userId,
			VideoID:   videoId,
			Content:   ContentText,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		user, err := model.GetUserById(userId)
		//添加评论
		model.AddComment(Comment)
		//添加完评论返回数据
		return entity.CommentResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "comment success"},
			Comment: entity.Comment{
				Id: Comment.ID,
				User: entity.User{
					Id:            userId,
					Name:          user.Name,
					FollowCount:   model.GetFolloweeCount(userId),
					FollowerCount: model.GetFollowerCount(userId),
				},
				Content:    ContentText,
				CreateDate: Comment.CreatedAt.Format("2006-01-02 15:04:05"),
			},
		}, nil

	} else if actionType == 2 {
		commentId := commentRequest.CommentId
		_, err := model.GetVideoById(videoId)
		_, err1 := model.GetUserById(userId)
		_, err2 := model.GetCommentById(commentId)
		if err != nil || err1 != nil || err2 != nil {
			return entity.CommentResponse{
				Response: entity.Response{StatusCode: 400, StatusMsg: "video is not exist or user is not exist or comment is not exist"},
				Comment:  entity.Comment{},
			}, err
		}
		model.DeleteCommentByCommentId(commentId)
		return entity.CommentResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "delete success"},
		}, nil

	}
	//actionType都不对返回错误
	return entity.CommentResponse{
		Response: entity.Response{StatusCode: 1, StatusMsg: "actionType is not right"},
		Comment:  entity.Comment{},
	}, nil
}

//遍历评论列表
func CommentList(videoId int64, userId int64) (entity.CommentListResponse, error) {
	//先判断ID是否存在
	_, err := model.GetVideoById(videoId)
	_, err1 := model.GetUserById(userId)
	if err != nil || err1 != nil {
		return entity.CommentListResponse{
			Response: entity.Response{StatusCode: 400, StatusMsg: "video is not exist or user is not exist"},
		}, err
	}
	commentList, err := model.GetCommentList(videoId, userId)
	if err != nil {
		return entity.CommentListResponse{Response: util.ServerErrorResponse}, err
	}
	if commentList == nil {
		return entity.CommentListResponse{Response: util.ListNilResponse}, err
	}
	return entity.CommentListResponse{
		Response:    util.SuccessResponse,
		CommentList: commentList,
	}, nil

}
