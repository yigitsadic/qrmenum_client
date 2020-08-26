package client

import (
	"testing"
)

func TestNewHTTPClient(t *testing.T) {
	t.Run("Should initialize as expected", func(t *testing.T) {
		url := "http://localhost:5000"
		got := NewHTTPClient(url)

		if got.BaseUrl != url {
			t.Errorf("Expected %s equal to %q", got.BaseUrl, url)
		}
	})
}
