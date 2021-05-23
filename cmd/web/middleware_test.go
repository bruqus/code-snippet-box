package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSecureHeaders(t *testing.T) {
	responseRecorder := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockNext := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	secureHeaders(mockNext).ServeHTTP(responseRecorder, req)
	result := responseRecorder.Result()

	frameOptions := result.Header.Get("X-Frame-Options")
	if frameOptions != "deny" {
		t.Errorf("expected %q, got %q", "deny", frameOptions)
	}

	xssProtection := result.Header.Get("X-XSS-Protection")
	if xssProtection != "1; mode=block" {
		t.Errorf("expected %q, got %q", "1; mode=block", xssProtection)
	}

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected %q, got %q", http.StatusOK, result.StatusCode)
	}

	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	if string(body) != "OK" {
		t.Errorf("expected body to equal %q", "OK")
	}
}
