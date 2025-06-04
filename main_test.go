package main

import (
	"os"
	"os/exec"
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
