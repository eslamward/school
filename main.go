package main

import (
	"os"
	"shool/cmd"
)

func main() {

	cmd.CliCommand().Run(os.Args)
}
