package routing

import (
	"fmt"
	"shool/config"

	"github.com/gin-gonic/gin"
)

func ServeRouting(router *gin.Engine) error {
	config := config.GetConfig()

	return router.Run(fmt.Sprintf("%v:%v", config.Server.Host, config.Server.Port))

}
