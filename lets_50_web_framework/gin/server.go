package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 获取一个默认路由 router
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"os":         runtime.GOOS,
			"go-version": runtime.Version(),
			"now_time":   time.Now().Format("2006-01-02 15:04:05"),
		})
	})
	//fmt.Println(runtime.GOOS)
	r.Run(":8080")
}
