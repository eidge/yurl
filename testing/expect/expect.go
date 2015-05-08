package expect

import (
	"testing"
)

func Equal(actual interface{}, expected interface{}, t *testing.T) {
	if actual != expected {
		t.Error("Expected %s, got %v", expected, actual)
	}
}

func Error(err error, t *testing.T) {
	if err == nil {
		t.Error("Expected error to not be nil.")
	}
}
