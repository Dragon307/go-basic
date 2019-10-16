package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/long_async", func(c *gin.Context) {
		// 创建在 goroutine 中使用的副本
		cCp := c.Copy()
		go func() {
			// 用 time.Sleep() 模拟一个长任务
			time.Sleep(5 * time.Second)

			// 这里使用的是复制上下文的 cCp
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)

		log.Println("Done ! in path " + c.Request.URL.Path)
	})

	r.Run(":8888")
}
