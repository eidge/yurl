/*
	Package main implements a simple CLI for the yurl tool.
*/
package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	//"log"
	"os"
)

func main() {
	cli.AppHelpTemplate = `{{.Name}} {{.Version}} - {{.Usage}}

  Usage:
     {{.Name}} FILENAME REQUEST_NAME [options]
  
  Example
     {{.Name}} users.yml get_users {{if len .Authors}}
  
  AUTHOR(S): 
     {{range .Authors}}{{ . }}{{end}}{{end}}
  
  COMMANDS:
     {{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
     {{end}}{{if .Flags}}{{end}}
`

	app := cli.NewApp()
	app.Name = "Yurl"
	app.Usage = "API requests made simple"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) {
		if len(c.Args()) == 0 {
			cli.ShowAppHelp(c)
		} else {
			config, err := parseYaml(c.Args()[0])
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			fmt.Printf("%v", config)
		}
	}
	app.Author = "Hugo Ribeira"
	app.Email = "hugoribeira@gmail.com"

	app.Commands = []cli.Command{
		{
			Name:    "glu",
			Aliases: []string{"a"},
			Action:  func(*cli.Context) { println("glu") },
		},
	}

	app.Run(os.Args)
}
