package controllers

import "github.com/gin-gonic/gin"

func Send(c *gin.Context) {

}

func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "hello",
	})
}