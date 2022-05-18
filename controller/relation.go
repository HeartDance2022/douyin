package controller

import (
	"douyin/entity"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {

	var relation entity.RelationRequest
	if err := c.ShouldBindJSON(&relation); err == nil {
		fmt.Printf("relation info:%#v\n", relation)
		if _, exist := usersLoginInfo[relation.Token]; exist {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
			_, err := service.Follow(&relation)
			if err != nil {
				return
			}
		} else {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		}
	}

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		followListResponse, err := service.GetFollowList(userId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK,
			followListResponse,
		)
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if _, exist := usersLoginInfo[token]; exist {
		followListResponse, err := service.GetFollowerList(userId)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK,
			followListResponse,
		)
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
