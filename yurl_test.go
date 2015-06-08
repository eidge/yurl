package main

import (
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"github.com/eidge/yurl/testing/expect"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// I don't really like this solution to do integration tests
// specially because it relies on yurl already beeing built.
// I haven't found a better solution yet, though.
var yurlCmd = "./yurl"

func startYurl(args ...string) *gexpect.ExpectSubprocess {
	child, err := gexpect.Spawn(yurlCmd + " " + strings.Join(args, " "))
	if err != nil {
		panic(err)
	}
	return child
}

func TestYurlWithNoArgs(t *testing.T) {
	yurl := startYurl()
	defer yurl.Close()

	// Expect help screen
	match, _ := yurl.ExpectRegex("Usage:")
	expect.Equal(t, match, true)
}

func TestYurlWithoutRequestName(t *testing.T) {
	yurl := startYurl("config/fixtures/simple.yml")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("Requests in")
	expect.Equal(t, match, true)
	match, _ = yurl.ExpectRegex("example_post")
	expect.Equal(t, match, true)
	match, _ = yurl.ExpectRegex("example_get")
	expect.Equal(t, match, true)
}

func TestYurlWithNonExistantFile(t *testing.T) {
	yurl := startYurl("non_existant_yaml request_name")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("no such file")
	expect.Equal(t, match, true)
}

func TestYurlWithNonExistantRequest(t *testing.T) {
	yurl := startYurl("config/fixtures/simple.yml request_name")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("not found")
	expect.Equal(t, match, true)
}

func TestYurlWithInvalidRequest(t *testing.T) {
	yurl := startYurl("config/fixtures/invalid_request.yml request_name")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("Invalid request")
	expect.Equal(t, match, true)
}

func TestYurlWithValidArguments(t *testing.T) {
	// Start test server
	server := httptest.NewUnstartedServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "it works")
		}))

	server.Listener, _ = net.Listen("tcp", ":8080")
	server.Start()
	defer server.Close()

	// Run yurl against the test server
	yurl := startYurl("config/fixtures/simple.yml integration_test")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("it works")
	expect.Equal(t, match, true)
}
