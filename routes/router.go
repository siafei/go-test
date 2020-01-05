package routes

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	router := gin.Default()
	registerWeb(router)

	api := router.Group("api")
	registerV1(api.Group("/v1"))
	return router
}
