package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

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
