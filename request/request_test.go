package request

import (
	"github.com/eidge/yurl/testing/expect"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidRequest(t *testing.T) {
	request := New()

	ok, err := request.IsValid()
	expect.Equal(t, ok, false)
	expect.Match(t, err.Error(), "Invalid request")
	expect.Match(t, err.Error(), "url cannot be blank")
	expect.Match(t, err.Error(), "method cannot be blank")
	expect.Match(t, err.Error(), "body_format cannot be blank")

	request.Method = "NotAllowed"
	ok, err = request.IsValid()
	expect.Match(t, err.Error(), "method is not allowed")

	request.BodyFormat = "NotAllowed"
	ok, err = request.IsValid()
	expect.Match(t, err.Error(), "body_format is not allowed")
}

func TestValidRequest(t *testing.T) {
	request := New()
	request.Url = "http://example.com"
	request.Method = "GET"
	request.BodyFormat = "raw"

	ok, err := request.IsValid()
	expect.Equal(t, ok, true)
	expect.NoError(t, err)
}

func TestRequestDoUsesConfigParameters(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		bodyBytes, _ := ioutil.ReadAll(r.Body)
		body := string(bodyBytes)

		expect.Equal(t, r.Method, "GET")
		expect.Equal(t, body, "Hi there!")
	}))
	defer ts.Close()

	request := New()
	request.Url = ts.URL
	request.Method = "GET"
	request.BodyFormat = "raw"
	request.Body = "Hi there!"

	_, err := request.Do()
	expect.NoError(t, err)
}
