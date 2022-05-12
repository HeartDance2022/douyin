package controller

import (
	"douyin/entity"
	"douyin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK, entity.UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, entity.UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}
