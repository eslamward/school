package routing

import (
	"shool/pkg/auth"

	"github.com/gin-gonic/gin"
)

func UserRouting(router *gin.Engine) {

	router.POST("/sign", auth.SignIn)
	router.GET("/sign", auth.SignForm)
	router.POST("/signup", auth.Register)
	router.POST("/signout", auth.Signout)
	router.POST("/reset", auth.ResetPassword)
	router.POST("/activate", auth.AcivateUser)
}
