package main

import "strings"

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

func resolveCommitType(input string) *CommitType {
	input = strings.ToLower(input)
	for _, ct := range commitTypes {
		if ct.Name == input || ct.Short == input {
			return &ct
		}
	}
	return nil
}
