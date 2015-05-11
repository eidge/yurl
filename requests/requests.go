/*
	Package requests is responsible for making http requests based on a Config object.
*/
package requests

import (
	"github.com/eidge/yurl/config"
	"net/http"
)

type Request struct{}

var client = &http.Client{}

func newRequest(config config.Request) (*http.Request, error) {
	request, err := http.NewRequest("GET", config.Url, nil)
	if err != nil {
		panic(err)
	}

	return request, nil
}
