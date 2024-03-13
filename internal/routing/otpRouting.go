package routing

import (
	"shool/pkg/otp"

	"github.com/gin-gonic/gin"
)

func OTPRouting(router *gin.Engine) {

	router.POST("/sendotp", otp.ResendOTP)
}
