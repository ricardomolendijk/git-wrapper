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
  --scope <scope>        Optional scope for the commit (e.g., api, parser)
  --breaking             Mark as a breaking change (adds ! after type/scope)
  --breaking-change <d>  Add a BREAKING CHANGE footer with description
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
    scope := ""
    breaking := false
    breakingDesc := ""
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
        case "--scope":
            i++
            if i < len(args) {
                scope = args[i]
            } else {
                fmt.Fprintln(os.Stderr, "Error: --scope requires a value")
                os.Exit(1)
            }
        case "--breaking":
            breaking = true
        case "--breaking-change":
            i++
            if i < len(args) {
                breakingDesc = args[i]
            } else {
                fmt.Fprintln(os.Stderr, "Error: --breaking-change requires a value")
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
        fmt.Fprintf(os.Stderr, "Warning: unknown commit type '%s', using chore.\n", commitTypeInput)
        // Fallback to a conventional, non-semver-impacting type
        commitType = &CommitType{"chore", "chore", "", "Chore"}
    }

    if hasMessage {
        // Disallow emojis in user-provided message
        if containsEmoji(message) {
            fmt.Fprintln(os.Stderr, "Error: emojis are not allowed in commit messages.")
            os.Exit(1)
        }
        // Conventional Commits: <type>: <description> with optional footer(s)
        header := commitType.Short
        if scope != "" {
            header = fmt.Sprintf("%s(%s)", header, scope)
        }
        if breaking {
            header += "!"
        }
        commitMsg := fmt.Sprintf("%s: %s", header, message)
        if ticket != "" {
            footer := formatTicketFooter(ticket)
            if footer != "" {
                commitMsg += "\n\n" + footer
            }
        }
        if breakingDesc != "" {
            commitMsg += "\n\nBREAKING CHANGE: " + breakingDesc
        }
        if containsEmoji(commitMsg) {
            fmt.Fprintln(os.Stderr, "Error: emojis are not allowed in commit messages.")
            os.Exit(1)
        }
        runCommand("git", append([]string{"commit", "-m", commitMsg}, passthrough...)...)
    } else {
        useEditorWithTemplate(commitType, ticket, scope, breaking, breakingDesc, passthrough)
    }
}
