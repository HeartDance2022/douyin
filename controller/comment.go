package controller

import (
	"douyin/entity"
	"douyin/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	commentText := c.Query("comment_text")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	//如果用户登录了并且存在
	if exist := service.GetLoginUser(token); exist != nil {
		userId := exist.ID
		commentResponse, err := service.Comment(
			entity.CommentRequest{
				Token:       token,
				VideoId:     videoId,
				ActionType:  int32(actionType),
				CommentText: commentText,
				CommentId:   commentId,
			}, userId,
		)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, commentResponse)
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	//如果用户登录了并且存在
	if exist := service.GetLoginUser(token); exist != nil {
		userId := exist.ID
		commentListResponse, err := service.CommentList(videoId, userId)
		if err != nil {
			fmt.Print(err)
		}
		c.JSON(http.StatusOK, commentListResponse)
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
