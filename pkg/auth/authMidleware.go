package auth

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func RegisterSession(router *gin.Engine) {

	store := cookie.NewStore([]byte("secret"))

	router.Use(sessions.Sessions("mysession", store))

}

func AuthMiddleWare(c *gin.Context) {
	session := sessions.Default(c)

	tok := session.Get("token")

	if tok == nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))

	}
	c.Next()

}
