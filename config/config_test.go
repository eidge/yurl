package config

import (
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func TestFromYaml(t *testing.T) {
	// Test file opening
	notExistentFile := "./FileThatDoesNotExist"
	_, err := FromYaml(notExistentFile)
	expect.Error(t, err)

	requestsFile := "./fixtures/yamlparser.yml"
	requests, err := FromYaml(requestsFile)
	if err != nil {
		t.Error("Yml parsing failed with error: %v", err)
	}

	// Globals
	expect.Equal(t, requests.Globals.BaseUrl, "https://example.com")
	expect.Equal(t, requests.Globals.Headers["AUTHENTICATION_TOKEN"], "123456")

	// Requests
	expect.Equal(t, len(requests.Requests), 2)

	expect.Equal(t, requests.Requests["example_post"].Path, "/example_endpoint")
	expect.Equal(t, requests.Requests["example_post"].Method, "POST")
	expect.Equal(t, requests.Requests["example_post"].BodyFormat, "json")

	expect.Equal(t, requests.Requests["example_post"].Body["first_name"], "John")
	expect.Equal(t, requests.Requests["example_post"].Body["last_name"], "Doe")

	expect.Equal(t, requests.Requests["example_post"].QueryString["format"], "json")
}
