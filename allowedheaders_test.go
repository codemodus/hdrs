package hedrs

import (
	"strings"
	"testing"
)

func TestAllowedHeadersAdd(t *testing.T) {
	h1 := Accept
	h2 := ContentType
	h3 := Date

	h := NewAllowedHeaders()
	mLen := len(h.hdrs)

	d := []struct {
		hdrs []string
		inc  int
	}{
		{[]string{h1, h2}, 2},
		{[]string{h2}, 0},
		{[]string{}, 0},
		{[]string{"", ""}, 0},
		{nil, 0},
		{[]string{h3}, 1},
	}

	for _, v := range d {
		h.Add(v.hdrs...)
		mLen += v.inc

		got := len(h.hdrs)
		want := mLen
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestAllowedHeadersRmv(t *testing.T) {
	h1 := Accept
	h2 := ContentType
	h3 := Date

	h := NewAllowedHeaders(h1, h2, h3)
	mLen := len(h.hdrs)

	d := []struct {
		hdrs []string
		inc  int
	}{
		{[]string{h1, h2}, -2},
		{[]string{}, 0},
		{[]string{"", ""}, 0},
		{nil, 0},
		{[]string{h2}, 0},
	}

	for _, v := range d {
		h.Rmv(v.hdrs...)
		mLen += v.inc

		got := len(h.hdrs)
		want := mLen
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestAllowedHeadersString(t *testing.T) {
	tFn := func(h *AllowedHeaders, stage string) {
		s := strings.Split(h.String(), ", ")

		got := len(s)
		want := len(h.hdrs)
		if got != want && got > 0 && s[0] != "" {
			t.Errorf("got %d, want %d @%s", got, want, stage)
		}

		for k := range h.hdrs {
			found := false
			for _, y := range s {
				if k == y {
					found = true
				}
			}

			if !found {
				t.Errorf("got %s, want %s @%s", "", k, stage)
			}
		}
	}

	h := NewAllowedHeaders()
	tFn(h, "newNone")

	h = NewAllowedHeaders(Accept)
	tFn(h, "newAccept")

	h.Add(ContentType, Date)
	tFn(h, "addContentTypeAndDate")

	h.Rmv(ContentType)
	tFn(h, "rmvContentType")
}
