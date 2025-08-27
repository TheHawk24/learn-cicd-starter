package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	apikey := http.Header{
		"Authorization": {"ApiKey 123245"},
	}

	var tests = []struct {
		name  string
		input http.Header
		want  string
	}{
		{"With ApiKey", apikey, "123245"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := GetAPIKey(test.input)
			if result == "" {
				t.Errorf("Failed test, got: %s, want: %s", result, test.want)
			}
		})
	}

}

func TestGetAPIKeyError(t *testing.T) {

	input := http.Header{
		"Authorization": {"123245"},
	}
	want := errors.New("malformed authorization header")
	_, err := GetAPIKey(input)
	if err == nil {
		t.Errorf("Failed test, got: %s, want: %v", err, want)
	}

}
