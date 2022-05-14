package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/makhmudovazeez/sms-router/middlewares"
)

func Route() {
	router := gin.Default()

	gin.Logger()

	v2 := router.Group("/v2/notice/sms", middlewares.Auth)
	{
		v2.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": ctx.Request.URL.Path,
			})
		})

		v2.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Hello",
			})
		})
	}

	router.Run(os.Getenv("PORT"))
}
