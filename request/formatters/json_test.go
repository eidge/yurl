package formatters

import (
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func TestJSONFormatter(t *testing.T) {
	var body interface{}
	body = map[string]interface{}{
		"id":   3,
		"name": "John Doe",
	}
	parser := new(JSONFormatter)

	bytes, err := parser.Parse(body)
	expect.NoError(t, err)
	expect.Equal(t, string(bytes), "{\"id\":3,\"name\":\"John Doe\"}")
}
