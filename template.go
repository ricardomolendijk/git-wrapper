package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func useEditorWithTemplate(commitType *CommitType, ticket, scope string, breaking bool, breakingDesc string, passthrough []string) {
    changelog := getChangelog()

    // Build optional ticket footer per Conventional Commits
    ticketFooter := ""
    if ticket != "" {
        if f := formatTicketFooter(ticket); f != "" {
            ticketFooter = "\n\n" + f
        }
    }

    // Build optional breaking change footer
    breakingFooter := ""
    if breakingDesc != "" {
        breakingFooter = "\n\nBREAKING CHANGE: " + breakingDesc
    }

    // Build the initial commit message template
    header := commitType.Short
    if scope != "" {
        header = fmt.Sprintf("%s(%s)", header, scope)
    }
    if breaking {
        header += "!"
    }
    template := fmt.Sprintf(`%s: EDIT DESCRIPTION

Description:
Explain what this commit does and why.

Changelog:
%s

Affected files:
%s%s%s
`, header, changelog, changelog, ticketFooter, breakingFooter)

	// Write the template to a temporary file
	tempFile := "/tmp/git_commit_template.txt"
    if err := os.WriteFile(tempFile, []byte(template), 0644); err != nil {
        fmt.Println("Error: failed to write commit template:", err)
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
        fmt.Println("Error: editor aborted:", err)
        os.Exit(1)
    }

	// Read and clean the commit message
	contentBytes, err := os.ReadFile(tempFile)
    if err != nil {
        fmt.Println("Error: failed to read commit message:", err)
        os.Exit(1)
    }
    content := removeCommentLines(string(contentBytes))
    if containsEmoji(content) {
        fmt.Println("Error: emojis are not allowed in commit messages.")
        os.Exit(1)
    }

	// Check if the title was edited
    if strings.Contains(content, "EDIT DESCRIPTION") {
        fmt.Println("Error: commit aborted: description was not edited.")
        os.Exit(1)
    }

	// Check if unchanged
    if content == original {
        fmt.Println("Warning: commit aborted: message was not edited.")
        os.Exit(1)
    }

	// Write the cleaned content back and commit
    if err := os.WriteFile(tempFile, []byte(content), 0644); err != nil {
        fmt.Println("Error: failed to write cleaned commit message:", err)
        os.Exit(1)
    }
	runCommand("git", append([]string{"commit", "--file", tempFile}, passthrough...)...)
	_ = os.Remove(tempFile)
}
