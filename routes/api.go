package routes

import (
	"github.com/gin-gonic/gin"
	"go-test/app/controller/api/demo"
)


func registerV1(router *gin.RouterGroup) {
	router.GET("/", demo.Index)
	router.POST("login",demo.Login)
	router.POST("register",demo.Register)
}

