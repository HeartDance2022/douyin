package controller

import (
	"douyin/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	t := c.Query("latest_time")
	fmt.Println(t)

	timestamp, err := strconv.Atoi(t)
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
	}

	if timestamp < 0 {
		timestamp = int(time.Now().Unix())
	}
	res, err := service.Feed(time.Unix(int64(timestamp)/1e3, 0))
	if err != nil {
		Fail(c, http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, res)
}
