package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/eidge/yurl/config"
	"github.com/eidge/yurl/request"
	"github.com/eidge/yurl/response_formatter"
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
	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return
	}

	// Read config file
	filename := c.Args()[0]
	requests, err := config.RequestsFromYaml(filename)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(-1)
	}

	// List requests if no request is provided
	if len(c.Args()) == 1 {
		fmt.Printf("Requests in %s:\n", filename)
		printRequestList(requests)
		return
	}

	// Find and make request
	requestName := c.Args()[1]
	request, ok := requests[requestName]
	if ok {
		resp, err := request.Do()
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(-1)
		}
		responseFormatter.Print(resp)
	} else {
		fmt.Printf("Request %s not found in %s\n", requestName, filename)
		os.Exit(-1)
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

func printRequestList(requests map[string]request.Request) {
	for requestName, request := range requests {
		fmt.Printf("  %s\t%s\n", request.Method, requestName)
	}
}
