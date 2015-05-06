package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type RequestsConfig struct {
	Globals  GlobalConfig
	Requests map[string]Request
}

type GlobalConfig struct {
	BaseUrl string "base_url"
	Headers map[string]string
}

type Request struct {
	Path        string
	Method      string
	BodyFormat  string            "body_format"
	Body        map[string]string // This should be an interface to respect yaml types!
	QueryString map[string]string "query_str"
}

/*
	parseYaml takes a path to the request yaml file and returns a map containing the
	the configuration for the requests defined in that file.
*/
func parseYaml(filename string) (config RequestsConfig, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	yaml.Unmarshal(data, &config)
	return
}
