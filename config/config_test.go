package config

import (
	"github.com/eidge/yurl/request"
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func readYaml(t *testing.T, filename string) map[string]request.Request {
	requestsFile := filename
	requests, err := RequestsFromYaml(requestsFile)
	if err != nil {
		t.Errorf("Yml parsing failed with error: %v", err)
		return requests
	}
	return requests
}

func TestRequestsFromYamlParsesYamlFile(t *testing.T) {
	// Test file opening
	notExistentFile := "./FileThatDoesNotExist"
	_, err := RequestsFromYaml(notExistentFile)
	expect.Error(t, err)

	requests := readYaml(t, "./fixtures/simple.yml")

	// Requests
	expect.Equal(t, len(requests), 5)

	expect.Equal(t, requests["example_post"].Path, "/example_endpoint")
	expect.Equal(t, requests["example_post"].Method, "POST")
	expect.Equal(t, requests["example_post"].BodyFormat, "json")

	body := requests["example_post"].Body.(map[interface{}]interface{})
	expect.Equal(t, body["first_name"], "John")
	expect.Equal(t, body["last_name"], "Doe")

	expect.Equal(t, requests["example_post"].QueryString["format"], "json")
}

func TestRequestsFromYamlReturnsErrorIfRequestIsInvalid(t *testing.T) {
	notExistentFile := "./fixtures/invalid_request.yml"
	_, err := RequestsFromYaml(notExistentFile)
	expect.Error(t, err)
}

func TestRequestsFromYamlSetsConfigurationFromGlobals(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	request := requests["example_post"]

	expect.Equal(t, request.BaseUrl, "http://localhost:8080")
	expect.DeepEqual(t, request.Headers, map[string]string{"AUTHENTICATION_TOKEN": "123456"})
}

func TestRequestsFromYamlSetsUrlFromBaseUrlAndPath(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	requestWithUrl := requests["example_get"]
	requestWithoutUrl := requests["example_post"]

	expect.Equal(t, requestWithUrl.Url, "https://test.example.com/example_endpoint_2")
	expect.Equal(t, requestWithoutUrl.Url, "http://localhost:8080/example_endpoint")
}

func TestHeadersAreMergedWithGlobalsHeaders(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	requestWithNoAuthToken := requests["example_get"]
	requestWithAuthToken := requests["override_auth_token"]

	expect.Equal(t, requestWithNoAuthToken.Headers["Accept-Encoding"], "application/json")
	expect.Equal(t, requestWithNoAuthToken.Headers["AUTHENTICATION_TOKEN"], "123456")

	expect.Equal(t, requestWithAuthToken.Headers["Accept-Encoding"], "application/json")
	expect.Equal(t, requestWithAuthToken.Headers["AUTHENTICATION_TOKEN"], "overriden!")
}

func TestSetDefaults(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	request := requests["example_with_no_method_or_body_format"]

	expect.Equal(t, request.BodyFormat, "json")
	expect.Equal(t, request.Method, "GET")
}
