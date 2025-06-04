package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

const testBin = "./git-wrapper-test-bin"

func TestMain(m *testing.M) {
	// Build the binary for CLI tests
	cmd := exec.Command("go", "build", "-o", testBin)
	if out, err := cmd.CombinedOutput(); err != nil {
		panic("build failed: " + string(out))
	}
	code := m.Run()
	_ = os.Remove(testBin)
	os.Exit(code)
}

func TestMain_Help(t *testing.T) {
	cmd := exec.Command(testBin, "--help")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if len(out) == 0 || !strings.Contains(string(out), "Usage") {
		t.Errorf("expected usage output, got %q", string(out))
	}
}

func TestMain_UnknownCommand(t *testing.T) {
	cmd := exec.Command(testBin, "unknown")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("expected error for unknown command")
	}
	if string(out) == "" || !strings.Contains(string(out), "Unknown command") {
		t.Errorf("expected unknown command error, got %q", string(out))
	}
}
