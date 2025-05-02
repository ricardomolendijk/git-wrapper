package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// CommitType represents a commit type with metadata
type CommitType struct {
	Name        string
	Short       string
	Emoji       string
	DisplayName string
}

var commitTypes = []CommitType{
	{"feature", "feat", "🚀", "Feature"},
	{"fix", "bug", "🐛", "Bugfix"},
	{"chore", "chore", "🔧", "Chore"},
	{"documentation", "docs", "📚", "Documentation"},
	{"refactor", "refactor", "♻️", "Refactor"},
	{"test", "test", "🧪", "Test"},
	{"perf", "perf", "⚡", "Performance"},
	{"ci", "ci", "📦", "CI/CD"},
	{"config", "cfg", "🔧", "Configuration Change"},
  {"network", "net", "🌐", "Network Change"},
	{"misc", "misc", "📝", "Miscellaneous"},
	{"first-commit", "first", "🏁", "First Commit"},
	{"milestone", "mile", "🏆", "Milestone"},
	{"release", "release", "🎯", "Release"},
	{"style", "style", "🎨", "Style Change"},
	{"revert", "revert", "⏪", "Revert"},
	{"merge", "merge", "🔀", "Merge"},
	{"security", "sec", "🔒", "Security Fix"},
	{"build", "build", "🏗️", "Build System"},
	{"deprecate", "depr", "🗑️", "Deprecation"},
	{"ux", "ux", "💡", "UX Improvement"},
	{"ui", "ui", "🖼️", "UI Update"},
	{"hotfix", "hotfix", "🚑", "Hotfix"},
	{"lint", "lint", "🧹", "Linting"},
	{"env", "env", "🌱", "Environment Setup"},
	{"legal", "legal", "📄", "Legal/Compliance"},
	{"infra", "infra", "🏭", "Infrastructure Change"},
	{"i18n", "intl", "🌍", "Internationalization"},
	{"analytics", "analytics", "📊", "Analytics/Tracking"},
	{"rollback", "rollback", "↩️", "Rollback"},
	{"prototype", "proto", "🧪", "Prototype/Experiment"},
	{"log", "log", "📝", "Logging"},
	{"monitoring", "mon", "📈", "Monitoring"},
	{"bump", "version", "🔖", "Version Bump"},
}

func main() {
	if len(os.Args) < 2 || os.Args[1] != "commit" {
		runCommand("git", os.Args[1:]...)
		return
	}

	args := os.Args[2:]
	ticket, commitTypeInput, message := "", "feature", ""
	hasMessage := false
	var passthrough []string

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "--ticket":
			i++
			if i < len(args) {
				ticket = args[i]
			}
		case "--type":
			i++
			if i < len(args) {
				commitTypeInput = args[i]
			}
		case "-m", "--message":
			i++
			if i < len(args) {
				message = args[i]
				hasMessage = true
			}
		default:
			passthrough = append(passthrough, args[i])
		}
	}

	commitType := resolveCommitType(commitTypeInput)
	if commitType == nil {
		fmt.Println("⚠️ Unknown commit type, using generic.")
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

func resolveCommitType(input string) *CommitType {
	input = strings.ToLower(input)
	for _, ct := range commitTypes {
		if ct.Name == input || ct.Short == input {
			return &ct
		}
	}
	return nil
}

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

func useEditorWithTemplate(commitType *CommitType, ticket string, passthrough []string) {
	changelog := getChangelog()

	// Build optional ticket section
	ticketBlock := ""
	if ticket != "" {
		ticketBlock = fmt.Sprintf("\n\n🔗 Related Ticket(s):\n%s", formatTicket2(ticket))
	}

	// Build the initial commit message template
	template := fmt.Sprintf(`%s %s: EDIT TITLE %s

📝 Description:
Explain what this commit does and why.

📦 Changelog:
%s

📁 Affected files:
%s%s
`, commitType.Emoji, commitType.DisplayName, formatTicket1(ticket), "- ...", changelog, ticketBlock)

	// Write the template to a temporary file
	tempFile := "/tmp/git_commit_template.txt"
	if err := os.WriteFile(tempFile, []byte(template), 0644); err != nil {
		fmt.Println("❌ Failed to write commit template:", err)
		os.Exit(1)
	}

	// Open the editor
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	original := strings.TrimSpace(template)
	edit := exec.Command(editor, "-n", tempFile)
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	edit.Stderr = os.Stderr
	if err := edit.Run(); err != nil {
		fmt.Println("❌ Editor aborted:", err)
		os.Exit(1)
	}

	// Read and clean the commit message
	contentBytes, err := os.ReadFile(tempFile)
	if err != nil {
		fmt.Println("❌ Failed to read commit message:", err)
		os.Exit(1)
	}
	content := removeCommentLines(string(contentBytes))

	// Check if the title was edited
	if strings.Contains(content, "EDIT TITLE") {
		fmt.Println("❌ Commit aborted: title was not edited.")
		os.Exit(1)
	}

	// Check if unchanged
	if content == original {
		fmt.Println("⚠️ Commit aborted: message was not edited.")
		os.Exit(1)
	}

	// Write the cleaned content back and commit
	if err := os.WriteFile(tempFile, []byte(content), 0644); err != nil {
		fmt.Println("❌ Failed to write cleaned commit message:", err)
		os.Exit(1)
	}
	runCommand("git", append([]string{"commit", "--file", tempFile}, passthrough...)...)
	_ = os.Remove(tempFile)
}
// formatTicket1 formats the ticket list as [#ticket] for each ticket.
func formatTicket1(ticket string) string {
	if ticket == "" {
		return ""
	}

	tickets := strings.Split(ticket, ",")
	var formattedTickets []string
	for _, t := range tickets {
		formattedTickets = append(formattedTickets, fmt.Sprintf("[%s]", strings.ToUpper(strings.TrimSpace(t))))
	}

	return strings.Join(formattedTickets, " ")
}

// formatTicket2 formats the ticket list as - #ticket for each ticket.
func formatTicket2(ticket string) string {
	if ticket == "" {
		return ""
	}

	tickets := strings.Split(ticket, ",")
	var formattedTickets []string
	for _, t := range tickets {
		formattedTickets = append(formattedTickets, fmt.Sprintf("- %s", strings.ToUpper(strings.TrimSpace(t))))
	}

	return strings.Join(formattedTickets, "\n")
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

func removeCommentLines(content string) string {
	var result []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
			// Add only lines that don't start with '#'
			if !strings.HasPrefix(strings.TrimSpace(line), "#") {
					result = append(result, line)
			}
	}

	return strings.Join(result, "\n")
}