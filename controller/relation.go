package controller

import (
	"douyin/entity"
	"douyin/service"
	"douyin/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 32)
	relation := entity.RelationRequest{Token: token, ActionType: int32(actionType), ToUserId: toUserId}
	if exist := service.GetLoginUser(relation.Token); exist != nil {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
		relation.UserId = exist.ID
		_, err := service.Follow(&relation)
		if err != nil {
			return
		}
	} else {
		c.JSON(http.StatusOK, util.TokenFailResponse)
	}

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if exist := service.GetLoginUser(token); exist != nil {
		curUserId := exist.ID
		followListResponse, err := service.GetFollowList(userId, curUserId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK,
			followListResponse,
		)
	} else {
		c.JSON(http.StatusOK, util.TokenFailResponse)
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if exist := service.GetLoginUser(token); exist != nil {
		curUserId := exist.ID
		followListResponse, err := service.GetFollowerList(userId, curUserId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK,
			followListResponse,
		)
	} else {
		c.JSON(http.StatusOK, util.TokenFailResponse)
	}
}
