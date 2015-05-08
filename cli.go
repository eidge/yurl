package main

import (
	"github.com/codegangsta/cli"
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
