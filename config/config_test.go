package config

import (
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func readYaml(t *testing.T, filename string) map[string]Request {
	requestsFile := filename
	requests, err := RequestsFromYaml(requestsFile)
	if err != nil {
		t.Error("Yml parsing failed with error: %v", err)
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
	expect.Equal(t, len(requests), 3)

	expect.Equal(t, requests["example_post"].Path, "/example_endpoint")
	expect.Equal(t, requests["example_post"].Method, "POST")
	expect.Equal(t, requests["example_post"].BodyFormat, "json")

	expect.Equal(t, requests["example_post"].Body["first_name"], "John")
	expect.Equal(t, requests["example_post"].Body["last_name"], "Doe")

	expect.Equal(t, requests["example_post"].QueryString["format"], "json")
}

func TestRequestsFromYamlSetsConfigurationFromGlobals(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	request := requests["example_post"]

	expect.Equal(t, request.BaseUrl, "https://example.com")
	expect.DeepEqual(t, request.Headers, map[string]string{"AUTHENTICATION_TOKEN": "123456"})
}

func TestRequestsFromYamlSetsUrlFromBaseUrlAndPath(t *testing.T) {
	requests := readYaml(t, "./fixtures/simple.yml")
	requestWithUrl := requests["example_get"]
	requestWithoutUrl := requests["example_post"]

	expect.Equal(t, requestWithUrl.Url, "https://test.example.com/example_endpoint_2")
	expect.Equal(t, requestWithoutUrl.Url, "https://example.com/example_endpoint")
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
