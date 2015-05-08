package config

import (
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func TestFromYaml(t *testing.T) {
	// Test file opening
	notExistentFile := "./FileThatDoesNotExist"
	_, err := FromYaml(notExistentFile)
	expect.Error(err, t)

	requestsFile := "./fixtures/yamlparser.yml"
	requests, err := FromYaml(requestsFile)
	if err != nil {
		t.Error("Yml parsing failed with error: %v", err)
	}

	// Globals
	expect.Equal(requests.Globals.BaseUrl, "https://example.com", t)
	expect.Equal(requests.Globals.Headers["AUTHENTICATION_TOKEN"], "123456", t)

	// Requests
	expect.Equal(len(requests.Requests), 2, t)

	expect.Equal(requests.Requests["example_post"].Path, "/example_endpoint", t)
	expect.Equal(requests.Requests["example_post"].Method, "POST", t)
	expect.Equal(requests.Requests["example_post"].BodyFormat, "json", t)

	expect.Equal(requests.Requests["example_post"].Body["first_name"], "John", t)
	expect.Equal(requests.Requests["example_post"].Body["last_name"], "Doe", t)

	expect.Equal(requests.Requests["example_post"].QueryString["format"], "json", t)
}
