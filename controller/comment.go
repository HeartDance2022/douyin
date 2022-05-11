package controller

import (
	"douyin/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, entity.CommentListResponse{
		Response:    entity.Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
