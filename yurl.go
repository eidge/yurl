/*
	Package main implements a simple CLI for the yurl tool.
*/
package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := newApp()

	// Define default action
	app.Action = actionMakeRequest

	// Define cli commands
	app.Commands = []cli.Command{
		{
			Name: "dummy",
			Action: func(c *cli.Context) {
			},
		},
	}

	app.Run(os.Args)
}
