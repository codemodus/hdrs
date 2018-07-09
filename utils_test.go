package hedrs

import (
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
