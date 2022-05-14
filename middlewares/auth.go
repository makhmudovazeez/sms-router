package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Token  string `header:"Authorization"`
	Domain string `header:"Domain"`
}

const BearerString = "Bearer "

func Auth(ctx *gin.Context) {
	header := Header{}

	if err := ctx.ShouldBindHeader(&header); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	_, err := header.splitToken()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"Error" : err.Error(),
		})
		return
	}

	ctx.Next()
}

func (h Header) splitToken() (string, error) {
	if !strings.Contains(h.Token, BearerString) && strings.Index(h.Token, BearerString) != 0 {
		return "", errors.New("Unauthorized error")
	}

	return h.Token[len(BearerString):], nil
}
