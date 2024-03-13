package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Signout(c *gin.Context) {

	session := sessions.Default(c)

	session.Clear()

	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Logout Successfully"})
}
