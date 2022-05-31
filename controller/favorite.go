package controller

import (
	"douyin/entity"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	VideoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	like := entity.LikeRequest{Token: token, VideoId: VideoId, ActionType: int32(actionType)}
	if exist := service.GetLoginUser(token); exist != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
		like.UserId = exist.ID
		_, err := service.Like(&like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.ServerErrorResponse)
		}
	} else {
		c.JSON(http.StatusOK, util.TokenFailResponse)
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if exist := service.GetLoginUser(token); exist != nil {
		videoList, err := service.GetFavoriteVideoList(userId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, util.ServerErrorResponse)
		}
		c.JSON(http.StatusOK,
			videoList,
		)
	} else {
		c.JSON(http.StatusOK, util.TokenFailResponse)
	}
}
