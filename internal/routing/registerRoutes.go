package routing

import (
	"shool/pkg/auth"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	auth.RegisterSession(router)
	ArticleRouting(router)
	UserRouting(router)
	OTPRouting(router)
}
