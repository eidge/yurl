/*
	Package config is responsible for reading and exposing request configuration variables
*/
package config

import (
	"github.com/eidge/yurl/request"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
)

type config struct {
	Globals  request.Request
	Requests map[string]request.Request
}

/*
	RequestsFromYaml takes a path to the request yaml file and returns a map containing the
	the configuration for the requests defined in that file.
*/
func RequestsFromYaml(filename string) (map[string]request.Request, error) {
	var config config
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(data, &config)

	for requestName, _ := range config.Requests {
		config.Requests[requestName] = applyGlobals(config.Requests[requestName], config.Globals)
		config.Requests[requestName] = mergeHeaders(config.Requests[requestName], config.Globals)
		config.Requests[requestName] = setUrlFromBaseUrl(config.Requests[requestName])
		config.Requests[requestName] = setDefaults(config.Requests[requestName])

		if valid, err := config.Requests[requestName].IsValid(); !valid {
			return nil, err
		}
	}

	return config.Requests, nil
}

// Apply values from globals to every zero-valued field in request.
func applyGlobals(request request.Request, globals request.Request) request.Request {
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

func setUrlFromBaseUrl(request request.Request) request.Request {
	if request.Url == "" {
		request.Url = request.BaseUrl + request.Path
	}
	return request
}

func mergeHeaders(request request.Request, globals request.Request) request.Request {
	for headerName, headerValue := range globals.Headers {
		if request.Headers[headerName] == "" {
			request.Headers[headerName] = headerValue
		}
	}
	return request
}

func setDefaults(request request.Request) request.Request {
	request.Method = firstNonZeroValue(request.Method, "GET").Interface().(string)
	request.BodyFormat = firstNonZeroValue(request.BodyFormat, "json").Interface().(string)
	return request
}
