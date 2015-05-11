package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/eidge/yurl/config"
	"os"
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Author = "Hugo Ribeira"
	app.Email = "hugoribeira@gmail.com"
	app.Name = "Yurl"
	app.Usage = "API requests made simple"
	app.Version = "0.1"

	cli.AppHelpTemplate = appHelpTemplate()
	return app
}

func actionMakeRequest(c *cli.Context) {
	if len(c.Args()) == 1 {
		fmt.Printf("No request provided. \n\n")
		cli.ShowAppHelp(c)
	} else if len(c.Args()) > 1 {
		// Read config file
		filename := c.Args()[0]
		requests, err := config.RequestsFromYaml(filename)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(-1)
		}

		// Find and make request
		requestName := c.Args()[1]
		requestConfig, ok := requests[requestName]
		if ok {
			fmt.Printf("%#v", requestConfig)
		} else {
			fmt.Printf("Request %s not found in %s\n", requestName, filename)
			os.Exit(-1)
		}
	} else {
		cli.ShowAppHelp(c)
	}
}

func appHelpTemplate() string {
	return `{{.Name}} {{.Version}} - {{.Usage}}

  Usage:
  	{{.Name}} FILENAME REQUEST_NAME [options]
  
  Example
  	{{.Name}} users.yml get_users {{if len .Authors}}
  
  AUTHOR(S): 
  	{{range .Authors}}{{ . }}{{end}}{{end}}
  
  COMMANDS:
  	{{range .Commands}}{{join .Names ", "}}{{ "\t" }}{{.Usage}}
  	{{end}}{{if .Flags}}{{end}}`
}
