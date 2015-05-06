package main

import (
	"testing"
)

func TestParseYaml(t *testing.T) {
	// Test file opening
	notExistentFile := "./FileThatDoesNotExist"
	_, err := parseYaml(notExistentFile)
	expectError(err, t)

	requestsFile := "./fixtures/yamlparser.yml"
	requests, err := parseYaml(requestsFile)
	if err != nil {
		t.Error("Yml parsing failed with error: %v", err)
	}

	// Globals
	expectEqual(requests.Globals.BaseUrl, "https://example.com", t)
	expectEqual(requests.Globals.Headers["AUTHENTICATION_TOKEN"], "123456", t)

	// Requests
	expectEqual(len(requests.Requests), 2, t)

	expectEqual(requests.Requests["example_post"].Path, "/example_endpoint", t)
	expectEqual(requests.Requests["example_post"].Method, "POST", t)
	expectEqual(requests.Requests["example_post"].BodyFormat, "json", t)

	expectEqual(requests.Requests["example_post"].Body["first_name"], "John", t)
	expectEqual(requests.Requests["example_post"].Body["last_name"], "Doe", t)

	expectEqual(requests.Requests["example_post"].QueryString["format"], "json", t)
}

func expectEqual(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Error("Expected %s, got %v", expected, actual)
	}
}

func expectError(err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error to not be nil.")
	}
}
