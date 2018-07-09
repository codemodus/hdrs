package hedrs

import (
	"net/http"
)

// Allower ...
type Allower interface {
	IsAllowed(string) bool
}

// Stringer ...
type Stringer interface {
	String() string
}

// CORSOrigins ...
func CORSOrigins(a Allower) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			o := r.Header.Get(Origin)

			if o == "" || !a.IsAllowed(o) {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Add(AccessControlAllowOrigin, o)

			next.ServeHTTP(w, r)
		})
	}
}

// CORSHeaders ...
func CORSHeaders(s Stringer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Add(AccessControlAllowHeaders, s.String())

			next.ServeHTTP(w, r)
		})
	}
}

// CORSMethods ...
func CORSMethods(s Stringer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodOptions {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Add(AccessControlAllowMethods, s.String())

			next.ServeHTTP(w, r)
		})
	}
}

// OptionsHalt ...
func OptionsHalt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
