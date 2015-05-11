/*
	Package config is responsible for reading and exposing request configuration variables
*/
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type config struct {
	Globals  Request
	Requests map[string]Request
}

type Request struct {
	BaseUrl     string "base_url"
	BodyFormat  string "body_format"
	Method      string
	Path        string
	Url         string
	Body        map[string]string // This should be an interface to respect yaml types!
	Headers     map[string]string
	QueryString map[string]string "query_str"
}

/*
	RequestsFromYaml takes a path to the request yaml file and returns a map containing the
	the configuration for the requests defined in that file.
*/
func RequestsFromYaml(filename string) (map[string]Request, error) {
	var config config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config.Requests, err
	}

	yaml.Unmarshal(data, &config)

	for requestName, _ := range config.Requests {
		config.Requests[requestName] = applyGlobals(config.Requests[requestName], config.Globals)
		config.Requests[requestName] = mergeHeaders(config.Requests[requestName], config.Globals)
		config.Requests[requestName] = setUrlFromBaseUrl(config.Requests[requestName])
	}

	return config.Requests, err
}

// Apply values from globals to every zero-valued field in request.
func applyGlobals(request Request, globals Request) Request {
	requestValue := reflect.ValueOf(&request).Elem()
	globalsValue := reflect.ValueOf(&globals).Elem()

	//Transverse all exported fields from Config struct
	for i := 0; i < requestValue.NumField(); i++ {
		requestField := requestValue.Field(i)
		globalsField := globalsValue.Field(i)
		newValue := firstNonZeroValue(requestField.Interface(), globalsField.Interface())

		requestField.Set(newValue)
	}

	return request
}

// firstNonBlank returns ValueOf(defaulValue) if value is a zero-value.
func firstNonZeroValue(value interface{}, defaultValue interface{}) reflect.Value {
	var returnValue interface{}

	switch value.(type) {
	case string:
		if value == "" {
			returnValue = defaultValue
		} else {
			returnValue = value
		}
	case map[string]string:
		if len(value.(map[string]string)) == 0 {
			returnValue = defaultValue
		} else {
			returnValue = value
		}
	default:
		panic("Type not supported")
	}
	return reflect.ValueOf(returnValue)
}

func setUrlFromBaseUrl(request Request) Request {
	if request.Url == "" {
		request.Url = request.BaseUrl + request.Path
	}
	return request
}

func mergeHeaders(request Request, globals Request) Request {
	for headerName, headerValue := range globals.Headers {
		if request.Headers[headerName] == "" {
			request.Headers[headerName] = headerValue
		}
	}
	return request
}
