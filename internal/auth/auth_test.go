package auth

import (
	"net/http"
	"testing"
)

func TestEmptyAPIKey(t *testing.T) {
	headers := http.Header{}
	actual, err := GetAPIKey(headers)

	if actual != "" || err == nil {
		t.Errorf("invalid input for testing Empty API Key")
	}
}

func TestInvalidAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "Bearer NotValid")

	actual, err := GetAPIKey(headers)

	if actual != "" || err == nil {
		t.Errorf("invalid input for testing Invalid API Key")
	}
}

func TestLackingAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey")

	actual, err := GetAPIKey(headers)

	if actual != "" || err == nil {
		t.Errorf("invalid input for testing Lacking API Key")
	}
}

func TestValidAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey some_keys")

	actual, err := GetAPIKey(headers)

	// intentional to test failure of CI
	if actual != "" || err != nil {
		t.Errorf("invalid input for testing Valid API Key")
	}
}
