package main

import (
	"github.com/gin-gonic/gin"
)

// 最初の設定処理
func init() {

}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}
