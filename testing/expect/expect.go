package expect

import (
	"testing"
)

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Error("Expected %s, got %v", expected, actual)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Error("Expected error to not be nil.")
	}
}
