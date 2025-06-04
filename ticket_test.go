package main

import "testing"

func TestFormatTicket1(t *testing.T) {
	cases := []struct {
		in   string
		out  string
	}{
		{"", ""},
		{"ABC-123", "[ABC-123]"},
		{"abc-123,def-456", "[ABC-123] [DEF-456]"},
		{"  xyz-1  ,  foo-2 ", "[XYZ-1] [FOO-2]"},
	}
	for _, c := range cases {
		got := formatTicket1(c.in)
		if got != c.out {
			t.Errorf("formatTicket1(%q) = %q; want %q", c.in, got, c.out)
		}
	}
}

func TestFormatTicket2(t *testing.T) {
	cases := []struct {
		in   string
		out  string
	}{
		{"", ""},
		{"ABC-123", "- ABC-123"},
		{"abc-123,def-456", "- ABC-123\n- DEF-456"},
		{"  xyz-1  ,  foo-2 ", "- XYZ-1\n- FOO-2"},
	}
	for _, c := range cases {
		got := formatTicket2(c.in)
		if got != c.out {
			t.Errorf("formatTicket2(%q) = %q; want %q", c.in, got, c.out)
		}
	}
}
