package morestrings

import "testing"

func testReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string

	} {
		{"Hello, world!", "!dlorw ,elloH"},
		{"", ""},
		{"abc", "cba"},

	}

	for _, c := range cases {
		got := ReverseRunes(c.in)
		if got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	
	}
	
}
