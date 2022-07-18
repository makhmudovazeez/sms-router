package router

import (
	"github.com/gin-gonic/gin"
	"github.com/makhmudovazeez/sms-router/internal/controllers"
)

func InitRoute() {
	router := gin.Default()

	sms := router.Group("v2/notice/sms")
	{
		sms.POST("/", controllers.Send)
		sms.GET("/", controllers.Get)
	}


	router.Run(":8080")
}