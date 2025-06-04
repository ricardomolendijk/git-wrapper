package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func getChangelog() string {
	cmd := exec.Command("git", "diff", "--cached", "--name-status")
	out, err := cmd.Output()
	if err != nil {
		return "- ..."
	}

	var result []string
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			status := parts[0]
			file := parts[1]
			prefix := map[string]string{"A": "+", "D": "-", "M": "~"}[status]
			result = append(result, fmt.Sprintf("%s %s", prefix, file))
		}
	}
	if len(result) == 0 {
		return "- ..."
	}
	return strings.Join(result, "\n")
}
