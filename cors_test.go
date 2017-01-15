package hedrs

import (
	"net/http"
	"net/http/httptest"
)
import "testing"

func TestCORS(t *testing.T) {
	d := []struct {
		origin string
		method string
		code   int
	}{
		{"", http.MethodConnect, http.StatusTeapot},
		{"", http.MethodDelete, http.StatusTeapot},
		{"", http.MethodGet, http.StatusTeapot},
		{"", http.MethodHead, http.StatusTeapot},
		{"", http.MethodOptions, http.StatusForbidden},
		{"http://test.example", http.MethodOptions, http.StatusTeapot},
		{"", http.MethodPatch, http.StatusTeapot},
		{"", http.MethodPost, http.StatusTeapot},
		{"", http.MethodPut, http.StatusTeapot},
		{"", http.MethodTrace, http.StatusTeapot},
	}

	ao := NewAllowedOrigins()
	ah := NewAllowedHeaders()
	mw := CORS(ao, ah)

	for _, v := range d {
		ao.Add(v.origin)

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
		}

		got := w.Header().Get(AccessControlAllowOrigin)
		want := v.origin
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
}
