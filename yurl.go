/*
	Package main implements a simple CLI for the yurl tool.
*/
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	//"log"
	"github.com/eidge/yurl/config"
	"os"
)

func main() {
	app := newApp()

	// Define default action
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
		} else {
			config, err := config.FromYaml(c.Args()[0])
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			fmt.Printf("%v", config)
		}
	}

	// Define cli commands
	app.Commands = []cli.Command{
		{
			Name:    "glu",
			Aliases: []string{"a"},
			Action:  func(*cli.Context) { println("glu") },
		},
	}

	app.Run(os.Args)
}
