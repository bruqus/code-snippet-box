package main

import (
	"bytes"
	"net/http"
	"testing"
)

func TestShowSnippet(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name         string
		urlPath      string
		expectedCode int
		expectedBody []byte
	}{
		{"Valid ID", "/snippet/1", http.StatusOK, []byte("Foo baz bar...")},
		{"Non-existent ID", "/snippet/2", http.StatusNotFound, nil},
		{"Negative ID", "/snippet/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/snippet/1.23", http.StatusNotFound, nil},
		{"String ID", "/snippet/foo", http.StatusNotFound, nil},
		{"Empty ID", "/snippet/", http.StatusNotFound, nil},
		{"Trailing slash", "/snippet/1/", http.StatusNotFound, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			code, _, body := ts.get(t, test.urlPath)

			if code != test.expectedCode {
				t.Errorf("expected %d, got %d", test.expectedCode, code)
			}

			if !bytes.Contains(body, test.expectedBody) {
				t.Errorf("expected body to contain %q", test.expectedBody)
			}
		})
	}
}