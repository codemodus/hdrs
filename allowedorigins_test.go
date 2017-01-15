package hedrs

import "testing"

func TestAllowedOriginsAdd(t *testing.T) {
	d1 := "http://test1.example"
	d2 := "http://test2.example"
	d3 := "http://test3.example"

	o := NewAllowedOrigins()
	mLen := len(o.hosts)

	d := []struct {
		doms []string
		inc  int
	}{
		{[]string{d1, d2}, 2},
		{[]string{d1}, 0},
		{[]string{}, 0},
		{[]string{"", ""}, 0},
		{nil, 0},
		{[]string{d3}, 1},
	}

	for _, v := range d {
		o.Add(v.doms...)
		mLen += v.inc

		got := len(o.hosts)
		want := mLen
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestAllowedOriginsRmv(t *testing.T) {
	d1 := "http://test1.example"
	d2 := "http://test2.example"
	d3 := "http://test3.example"

	o := NewAllowedOrigins(d1, d2, d3)
	mLen := len(o.hosts)

	d := []struct {
		doms []string
		inc  int
	}{
		{[]string{d1, d2}, -2},
		{[]string{}, 0},
		{[]string{"", ""}, 0},
		{nil, 0},
		{[]string{d1}, 0},
	}

	for _, v := range d {
		o.Rmv(v.doms...)
		mLen += v.inc

		got := len(o.hosts)
		want := mLen
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}

func TestAllowedOriginsIsAllowed(t *testing.T) {
	dom1 := "http://test1.example:12345"
	dom2 := "https://www.test2.example"

	o := NewAllowedOrigins(dom1, dom2)

	d := []struct {
		dom string
		res bool
	}{
		{"http://test1.example", true},
		{"https://test1.example", false},
		{"http://www.test1.example", false},
		{"https://www.test2.example", true},
		{"https://test2.example", false},
		{"", false},
	}

	for _, v := range d {
		got := o.IsAllowed(v.dom)
		want := v.res
		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	}
}
