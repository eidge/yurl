package request

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/eidge/yurl/request/formatters"
	"net/http"
)

var Formatters = make(map[string]Formatter)

type Formatter interface {
	Parse(interface{}) ([]byte, error)
}

func init() {
	Formatters["raw"] = new(formatters.RawFormatter)
	Formatters["json"] = new(formatters.JSONFormatter)
}

var allowedRequestTypes = []string{"GET", "POST", "PATCH", "HEAD", "PUT", "DELETE"}

type Request struct {
	BaseUrl     string "base_url"
	Path        string
	BodyFormat  string "body_format"
	Method      string
	Url         string
	Body        interface{}
	Headers     map[string]string
	QueryString map[string]string "query_str"
}

func New() *Request {
	return new(Request)
}

func (request *Request) Do() (*http.Response, error) {
	if valid, err := request.IsValid(); !valid {
		return nil, err
	}

	req, err := request.newHttpRequest()
	if err != nil {
		return nil, errors.New("Could not create request: " + err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("Could not complete request: " + err.Error())
	}

	return resp, nil
}

func (request Request) IsValid() (bool, error) {
	err := newValidationError()

	if request.Url == "" {
		err.addExplanation("url", "cannot be blank")
	}
	if request.Method == "" {
		err.addExplanation("method", "cannot be blank")
	} else {
		if !isStringInArray(request.Method, allowedRequestTypes) {
			err.addExplanation("method", "is not allowed")
		}
	}
	if request.BodyFormat == "" {
		err.addExplanation("body_format", "cannot be blank")
	} else {
		if _, ok := Formatters[request.BodyFormat]; !ok {
			err.addExplanation("body_format", "is not allowed")
		}
	}

	if len(err.explanations) == 0 {
		return true, nil
	} else {
		return false, err
	}
}

func (request *Request) newHttpRequest() (*http.Request, error) {
	body, err := request.parseBody()
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(request.Method, request.Url, body)
	if err != nil {
		return nil, err
	}

	for k, v := range request.Headers {
		httpReq.Header.Set(k, v)
	}

	return httpReq, nil
}

func (request *Request) parseBody() (*bytes.Buffer, error) {
	var parser Formatter
	var ok bool

	if parser, ok = Formatters[request.BodyFormat]; !ok {
		return nil, errors.New(fmt.Sprintf("No parse for format: %s", request.BodyFormat))
	}

	body, err := parser.Parse(request.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(body), nil
}
