package testutils

import "testing"

func ExpectToMatchInt(t *testing.T, actual int, expected int) {
	if actual != expected {
		t.Errorf("got %d, wanted %d", actual, expected)
	}
}

func ExpectToMatchString(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("got %s, wanted %s", actual, expected)
	}
}

func ExpectToMatchBool(t *testing.T, actual bool, expected bool) {
	if actual != expected {
		t.Errorf("got %t, wanted %t", actual, expected)
	}
}

func ExpectToErrorNil(t *testing.T, actual error) {
	if actual != nil {
		t.Errorf("did not expected error, got %e", actual)
	}
}

func ExpectToBeNil(t *testing.T, actual any) {
	if actual != nil {
		t.Errorf("was not nil, got %e", actual)
	}
}
