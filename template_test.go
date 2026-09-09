package main

import (
	"testing"
)

func TestBuildCommitTemplate_NoTicket(t *testing.T) {
    // This is a smoke test for template generation logic (not for editor interaction)
    ct := &CommitType{"feature", "feat", "", "Feature"}
    ticket := ""
    changelog := "- ..."
    affected := "changed.go"
    ticketBlock := ""

    template := buildCommitTemplate(ct, ticket, changelog, affected, ticketBlock)
    if template == "" || template == "EDIT DESCRIPTION" {
        t.Error("template should not be empty or default")
    }
}

func buildCommitTemplate(commitType *CommitType, ticket, changelog, affected, ticketBlock string) string {
    base := commitType.Short + ": EDIT DESCRIPTION\n\n" +
        "Description:\nExplain what this commit does and why.\n\n" +
        "Changelog:\n" + changelog + "\n\n" +
        "Affected files:\n" + affected
    return base + ticketBlock + "\n"
}
