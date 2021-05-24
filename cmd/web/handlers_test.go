package main

import (
	"bytes"
	"net/http"
	"net/url"
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
func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/user/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name         string
		userName     string
		userEmail    string
		userPassword string
		csrfToken    string
		expectedCode int
		expectedBody []byte
	}{
		{"Valid submission", "Keks", "test@example.com", "validPa$$word", csrfToken, http.StatusSeeOther, nil},
		{"Empty name", "", "test@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty email", "Keks", "", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty password", "Keks", "test@example.com", "", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Invalid email (incomplete domain)", "Keks", "test@example.", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", "Keks", "testexample.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing local part)", "Keks", "@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", "Keks", "test@example.com", "pas$$", csrfToken, http.StatusOK, []byte("This field is too short (minimum is 6 characters)")},
		{"Duplicate email", "Keks", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "", "wrongToken", http.StatusBadRequest, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("name", test.userName)
			form.Add("email", test.userEmail)
			form.Add("password", test.userPassword)
			form.Add("csrf_token", test.csrfToken)

			code, _, body := ts.postForm(t, "/user/signup", form)

			if code != test.expectedCode {
				t.Errorf("expected %d; got %d", test.expectedCode, code)
			}

			if !bytes.Contains(body, test.expectedBody) {
				t.Errorf("expected body %s to contain %q", body, test.expectedBody)
			}
		})
	}
}
