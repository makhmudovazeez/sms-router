package controllers

import "github.com/gin-gonic/gin"

type SMSInterface interface {
	Show(*gin.Context)
}

func send()