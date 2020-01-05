package routes

import "github.com/gin-gonic/gin"

func registerWeb(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "web home",
		})
	})

}

