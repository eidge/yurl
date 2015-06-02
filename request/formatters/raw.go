package formatters

import (
	"errors"
)

type RawFormatter struct{}

func (*RawFormatter) Parse(body interface{}) ([]byte, error) {
	if body, ok := body.(string); ok {
		return []byte(body), nil
	} else {
		return nil, errors.New("Could not parse body as a string.")
	}
}
