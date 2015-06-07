package responseFormatter

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"text/template"
)

type responseOutput struct {
	Protocol string
	Status   string
	Headers  string
	Body     string
}

var outputTemplate = `
{{.Protocol}} {{.Status}}

{{.Headers}}

{{.Body}}
`

func Print(response *http.Response) {
	tmpl, err := template.New("response").Parse(outputTemplate)
	if err != nil {
		panic(err.Error())
	}

	rspOutput := responseOutput{
		response.Proto,
		response.Status,
		formatHeadersMap(response.Header),
		formatBody(response.Body),
	}

	err = tmpl.Execute(os.Stdout, rspOutput)
	if err != nil {
		panic("Could not format response: " + err.Error())
	}
}

func formatHeadersMap(headers http.Header) string {
	headerKeys := sortedHeaderKeys(headers)

	headerStr := ""
	for _, k := range headerKeys {
		headerStr += fmt.Sprintf("%s: %s\n", k, headers.Get(k))
	}

	return headerStr
}

func formatBody(body io.ReadCloser) string {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		panic("Could not read body: " + err.Error())
	}
	return string(bodyBytes)
}

// Helpers

func sortedHeaderKeys(headers http.Header) []string {
	headerKeys := []string{}

	for key := range headers {
		headerKeys = append(headerKeys, key)
	}
	sort.Strings(headerKeys)

	return headerKeys
}
