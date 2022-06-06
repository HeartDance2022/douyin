package main

import (
	"douyin/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	//初始化连接池
	dao.InitRedis()
	r := gin.Default()

	initRouter(r)

	err := dao.InitMySQL()
	if err != nil {
		return
	}
	err = r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
