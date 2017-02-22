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
		origin string
		method string
		code   int
	}{
		{esite, http.MethodConnect, http.StatusTeapot},
		{esite, http.MethodDelete, http.StatusTeapot},
		{esite, http.MethodGet, http.StatusTeapot},
		{esite, http.MethodHead, http.StatusTeapot},
		{esite, http.MethodOptions, http.StatusTeapot},
		{esite, http.MethodPatch, http.StatusTeapot},
		{esite, http.MethodPost, http.StatusTeapot},
		{esite, http.MethodPut, http.StatusTeapot},
		{esite, http.MethodTrace, http.StatusTeapot},

		{bsite, http.MethodConnect, http.StatusForbidden},
		{bsite, http.MethodDelete, http.StatusForbidden},
		{bsite, http.MethodGet, http.StatusForbidden},
		{bsite, http.MethodHead, http.StatusForbidden},
		{bsite, http.MethodOptions, http.StatusForbidden},
		{bsite, http.MethodPatch, http.StatusForbidden},
		{bsite, http.MethodPost, http.StatusForbidden},
		{bsite, http.MethodPut, http.StatusForbidden},
		{bsite, http.MethodTrace, http.StatusForbidden},
	}

	ao := NewAllowedOrigins()
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

		gotCode := w.Code
		wantCode := v.code
		if gotCode != wantCode {
			t.Errorf("got %d, want %d", gotCode, wantCode)
			continue
		}

		if wantCode == http.StatusForbidden {
			continue
		}

		got := w.Header().Get(AccessControlAllowOrigin)
		want := esite
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

	ah := NewAllowedHeaders(ehdr)
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
