package router

import (
	"mjs/config"
	"mjs/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": config.Mysqldb,
		})
	})

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	user := r.Group("/user")
	{
		user.GET("/info/:userId", controllers.GetUserInfo)
	}
	notice := r.Group("/notice")
	{
		notice.POST("/list", controllers.GetNotices)
	}

	return r
}
