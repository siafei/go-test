package routes

import (
	"github.com/gin-gonic/gin"
	"go-test/app/middleware"
	"go-test/bootstrap"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.NoMethod(bootstrap.HandleNotFound)
	router.NoRoute(bootstrap.HandleNotFound)

	registerWeb(router)

	api := router.Group("api")

	//引用全局错误处理中间件
	api.Use(middleware.ErrHandler())
	
	registerV1(api.Group("/v1"))
	return router
}
