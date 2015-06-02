package expect

import (
	"reflect"
	"regexp"
	"testing"
)

func Equal(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Errorf("Expected %s, got %v", expected, actual)
	}
}

func DeepEqual(t *testing.T, actual interface{}, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, got %v", expected, actual)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to not be nil.")
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected error to be nil. Got: %s", err)
	}
}

func Match(t *testing.T, text string, regex string) {
	match, _ := regexp.MatchString(regex, text)
	if !match {
		t.Errorf("Expected \"%s\" to match /%s/, but it did not.", text, regex)
	}
}
