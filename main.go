package main

import (
	"douyin/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	initRouter(r)

	err := dao.InitMySQL()
	if err != nil {
		return
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
