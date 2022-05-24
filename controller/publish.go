package controller

import (
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, util.ServerErrorResponse)
	}
	c.JSON(http.StatusOK, service.PostVideo(form))
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userId := c.Query("user_id")
	token := c.Query("token")
	c.JSON(http.StatusOK, service.PostList(userId, token))
}
