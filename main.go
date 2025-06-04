package main

import (
	"fmt"
	"os"
)

const (
	defaultCommitType = "feature"
	version           = "1.0.0"
)

func printUsage() {
	fmt.Println(`Usage: git-wrapper commit [options]

Options:
  --ticket <ticket>      Add ticket reference(s) (comma-separated)
  --type <type>          Commit type (feature, fix, chore, ...)
  -m, --message <msg>    Commit message
  -h, --help             Show this help message
  --version              Show version
  [other git args]       Passed through to git`)
}

func printVersion() {
	fmt.Printf("git-wrapper version %s\n", version)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] != "commit" {
		runCommand("git", os.Args[1:]...)
		return
	}

	args := os.Args[2:]
	ticket, commitTypeInput, message := "", defaultCommitType, ""
	hasMessage := false
	var passthrough []string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--ticket":
			i++
			if i < len(args) {
				ticket = args[i]
			} else {
				fmt.Fprintln(os.Stderr, "Error: --ticket requires a value")
				os.Exit(1)
			}
		case "--type":
			i++
			if i < len(args) {
				commitTypeInput = args[i]
			} else {
				fmt.Fprintln(os.Stderr, "Error: --type requires a value")
				os.Exit(1)
			}
		case "-m", "--message":
			i++
			if i < len(args) {
				message = args[i]
				hasMessage = true
			} else {
				fmt.Fprintln(os.Stderr, "Error: -m/--message requires a value")
				os.Exit(1)
			}
		case "-h", "--help":
			printUsage()
			os.Exit(0)
		case "--version":
			printVersion()
			os.Exit(0)
		default:
			passthrough = append(passthrough, args[i])
		}
	}

	commitType := resolveCommitType(commitTypeInput)
	if commitType == nil {
		fmt.Fprintf(os.Stderr, "⚠️ Unknown commit type '%s', using generic.\n", commitTypeInput)
		commitType = &CommitType{"misc", "misc", "📝", "Miscellaneous"}
	}

	if hasMessage {
		commitMsg := fmt.Sprintf("%s %s: %s", commitType.Emoji, commitType.DisplayName, message)
		if ticket != "" {
			commitMsg += fmt.Sprintf(" [%s]", ticket)
		}
		runCommand("git", append([]string{"commit", "-m", commitMsg}, passthrough...)...)
	} else {
		useEditorWithTemplate(commitType, ticket, passthrough)
	}
}
