package main

import (
	"testing"
)

func TestGetChangelog_Empty(t *testing.T) {
	// This test will pass if there are no staged changes or not in a git repo
	got := getChangelog()
	if got != "- ..." && got == "" {
		t.Errorf("getChangelog() = %q; want '- ...' or non-empty string", got)
	}
}
