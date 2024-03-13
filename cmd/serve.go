package cmd

import (
	"shool/config"
	"shool/internal/database"
	"shool/internal/routing"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func CliCommand() *cli.App {
	return &cli.App{
		Name:        "School App",
		Description: "School app to manage and control The School",
		Commands: []cli.Command{
			{
				Name:        "Serve",
				Description: "Serve The App With Host And Port",
				Aliases:     []string{"serve", "S", "s"},
				Action: func(c *cli.Context) error {

					return Serve()
				},
			},
		},
	}
}

func Serve() error {
	config.SetConfig()
	database.ConnectToDatabase()
	router := gin.Default()
	router.LoadHTMLGlob("./internal/html/*.html")

	routing.Register(router)
	return routing.ServeRouting(router)

}
