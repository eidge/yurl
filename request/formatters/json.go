package formatters

import (
	"encoding/json"
)

type JSONFormatter struct{}

func (JSONFormatter) Parse(body interface{}) ([]byte, error) {
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}
