package main

import (
	"github.com/ThomasRooney/gexpect"
	"github.com/eidge/yurl/testing/expect"
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
	yurl := startYurl("some_yaml")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("No request provided")
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

func TestYurlWithValidArguments(t *testing.T) {
	yurl := startYurl("config/fixtures/simple.yml example_post")
	defer yurl.Close()

	match, _ := yurl.ExpectRegex("config.Request") // FIXME: Placeholder
	expect.Equal(t, match, true)
}
