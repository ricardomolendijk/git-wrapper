package main

import "testing"

func TestResolveCommitType(t *testing.T) {
	cases := []struct {
		in   string
		out  string // expected DisplayName
		found bool
	}{
		{"feature", "Feature", true},
		{"feat", "Feature", true},
		{"FIX", "Bugfix", true},
		{"unknown", "", false},
		{"chore", "Chore", true},
		{"docs", "Documentation", true},
	}
	for _, c := range cases {
		ct := resolveCommitType(c.in)
		if c.found && (ct == nil || ct.DisplayName != c.out) {
			t.Errorf("resolveCommitType(%q) = %v; want DisplayName %q", c.in, ct, c.out)
		}
		if !c.found && ct != nil {
			t.Errorf("resolveCommitType(%q) = %v; want nil", c.in, ct)
		}
	}
}
