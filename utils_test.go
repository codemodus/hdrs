package hedrs

import (
	"strings"
	"testing"
)

func TestPrunePort(t *testing.T) {
	d := []struct {
		in  string
		out string
	}{
		{"https://localhost:4040", "https://localhost"},
		{"https://localhost", "https://localhost"},
		{"http://localhost:4040", "http://localhost"},
		{"http://localhost", "http://localhost"},
		{"localhost:4040", "localhost"},
		{"localhost", "localhost"},
		{"https://127.0.0.1:4040", "https://127.0.0.1"},
		{"https://127.0.0.1", "https://127.0.0.1"},
		{"http://127.0.0.1:4040", "http://127.0.0.1"},
		{"http://127.0.0.1", "http://127.0.0.1"},
		{"127.0.0.1:4040", "127.0.0.1"},
		{"127.0.0.1", "127.0.0.1"},
	}

	for _, v := range d {
		got := prunePort(v.in)
		want := v.out
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	}
}

func TestKeysToString(t *testing.T) {
	as := struct{}{}

	d := []struct {
		m map[string]struct{}
	}{
		{
			map[string]struct{}{},
		},
		{
			map[string]struct{}{
				"test": as,
				"nest": as,
				"zest": as,
			},
		},
		{
			map[string]struct{}{
				"test": as,
				"nest": as,
			},
		},
	}

	for _, v := range d {
		s := strings.Split(keysToString(v.m), ", ")

		got := len(s)
		want := len(v.m)
		if got != want && got > 0 && s[0] != "" {
			t.Errorf("got %d, want %d for %v", got, want, v.m)
		}

		for k := range v.m {
			found := false
			for _, y := range s {
				if k == y {
					found = true
				}
			}

			if !found {
				t.Errorf("got %s, want %s in %s", "", k, v.m)
			}
		}

	}
}
