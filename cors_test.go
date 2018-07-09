package hedrs

import (
	"net/http"
	"net/http/httptest"
)
import "testing"

func TestCORSOrigins(t *testing.T) {
	esite := "https://example.com"
	bsite := "http://bad.example"

	d := []struct {
		origin  string
		method  string
		allowed string
	}{
		{esite, http.MethodConnect, esite},
		{esite, http.MethodDelete, esite},
		{esite, http.MethodGet, esite},
		{esite, http.MethodHead, esite},
		{esite, http.MethodOptions, esite},
		{esite, http.MethodPatch, esite},
		{esite, http.MethodPost, esite},
		{esite, http.MethodPut, esite},
		{esite, http.MethodTrace, esite},

		{bsite, http.MethodConnect, ""},
		{bsite, http.MethodDelete, ""},
		{bsite, http.MethodGet, ""},
		{bsite, http.MethodHead, ""},
		{bsite, http.MethodOptions, ""},
		{bsite, http.MethodPatch, ""},
		{bsite, http.MethodPost, ""},
		{bsite, http.MethodPut, ""},
		{bsite, http.MethodTrace, ""},
	}

	ao := NewAllowedRegistry()
	mw := CORSOrigins(ao)

	for _, v := range d {
		ao.Add(esite)

		w := httptest.NewRecorder()
		r, err := http.NewRequest(v.method, "/", nil)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		r.Header.Set(Origin, v.origin)

		mw(http.RedirectHandler("/", http.StatusTeapot)).ServeHTTP(w, r)

		got := w.Header().Get(AccessControlAllowOrigin)
		want := v.allowed
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}

func TestCORSHeaders(t *testing.T) {
	ehdr := "test"

	d := []struct {
		method  string
		headers string
	}{
		{http.MethodConnect, ""},
		{http.MethodDelete, ""},
		{http.MethodGet, ""},
		{http.MethodHead, ""},
		{http.MethodOptions, ehdr},
		{http.MethodPatch, ""},
		{http.MethodPost, ""},
		{http.MethodPut, ""},
		{http.MethodTrace, ""},
	}

	ah := NewValues(ehdr)
	mw := CORSHeaders(ah)

	for _, v := range d {
		w := httptest.NewRecorder()
		r, err := http.NewRequest(v.method, "/", nil)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		mw(http.RedirectHandler("/", http.StatusTeapot)).ServeHTTP(w, r)

		got := w.Header().Get(AccessControlAllowHeaders)
		want := v.headers
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
