package main

import "testing"

func TestRemoveCommentLines(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"# comment\nfoo\nbar\n# another\n", "foo\nbar"},
		{"foo\n#comment\n\nbar\n", "foo\nbar"},
		{"\n#c\n\n", ""},
		{"foo\n   # indented\nbar\n", "foo\nbar"},
		{"foo\n\nbar\n", "foo\nbar"},
	}
	for _, c := range cases {
		got := removeCommentLines(c.in)
		if got != c.out {
			t.Errorf("removeCommentLines(%q) = %q; want %q", c.in, got, c.out)
		}
	}
}
