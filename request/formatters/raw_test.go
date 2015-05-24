package formatters

import (
	"github.com/eidge/yurl/testing/expect"
	"testing"
)

func TestRawFormatter(t *testing.T) {
	var body interface{}
	body = "Hi there!"
	parser := new(RawFormatter)

	bytes, err := parser.Parse(body)
	expect.NoError(t, err)
	expect.Equal(t, string(bytes), "Hi there!")
}
