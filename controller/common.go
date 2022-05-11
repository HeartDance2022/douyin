package controller

import (
	"douyin/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Fail(ctx *gin.Context, httpStatus int, err error) {
	ctx.JSON(http.StatusOK, entity.FeedResponse{
		Response: entity.Response{StatusCode: 400, StatusMsg: err.Error()},
	})
}
