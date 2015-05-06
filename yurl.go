/*
	Package main implements a simple CLI for the yurl tool.
*/
package main

import (
	//"fmt"
	"github.com/codegangsta/cli"
	//"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Yurl"
	app.Usage = "API requests made simple"
	app.Version = "0.01"

	app.Run(os.Args)
}
