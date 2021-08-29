package employee

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerForbidden(t *testing.T) {
	// build request
	r := httptest.NewRequest(http.MethodGet, "/employee?phone=1111111111", nil)
	w := httptest.NewRecorder()

	// test handler
	SearchHandler(w, r)

	// validate response
	res := w.Result()
	if res.StatusCode != http.StatusForbidden {
		t.Fatal(res.StatusCode)
	}
}

func TestHandlerBadRequest(t *testing.T) {
	// build request
	r := httptest.NewRequest(http.MethodGet, "/employee", nil)
	w := httptest.NewRecorder()

	// add header
	r.Header.Add("GovernMint-token", "pa$$word")

	// test handler
	SearchHandler(w, r)

	// validate response
	res := w.Result()
	if res.StatusCode != http.StatusBadRequest {
		t.Fatal(res.StatusCode)
	}
}

func TestHandlerOK(t *testing.T) {
	// build request
	r := httptest.NewRequest(http.MethodGet, "/employee?phone=1111111111", nil)
	w := httptest.NewRecorder()

	// add header
	r.Header.Add("GovernMint-token", "pa$$word")

	// test handler
	SearchHandler(w, r)

	// validate response
	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatal(res.StatusCode)
	}
}
