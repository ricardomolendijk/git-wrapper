# Git Commit Helper

🚀 A customizable Git commit wrapper that enforces structured commit messages using emojis, ticket references, and changelogs.

---

## ✨ Features

- Supports multiple commit types (Feature, Bugfix, Documentation, etc.)
- Adds emoji, changelog, and optional ticket references
- Opens your preferred editor for rich commit messages
- Supports `-m` inline message or editor-based commit templates
- Falls back to standard `git` behavior for non-commit commands

---

## 📦 Installation

`go build -o git-wrapper main.go`
`mv git-wrapper ~/bin/git-wrapper`

make sure you set an `alias git='~/bin/git-wrapper'` in your `~/.zshrc or ~/.bashrc`

🔧 Usage
Run it just like git, but use the commit subcommand for enhanced functionality:

`git commit --type feat --ticket ABC-123 -m "add user registration"`

Or use the editor for detailed commit messages:

`git commit --type fix --ticket DEF-456`
Optional Flags
Flag Description
`--type` Commit type (feat, fix, docs, etc.)
`--ticket` Ticket ID(s), comma-separated
`-m`, `--message` Commit message (if skipping editor)

🧠 Commit Message Template (Editor)
When using the editor, a rich template is provided:

```
🚀 Feature: EDIT TITLE [ABC-123]

📝 Description:
Explain what this commit does and why.

📦 Changelog:

- added_file.go
  ~ modified_file.go

📁 Affected files:

- added_file.go
  ~ modified_file.go

🔗 Related Ticket(s):

- ABC-123
```

If the title still contains EDIT TITLE, the commit will be aborted.

## ✅ Supported Commit Types

| Name          | Short     | Emoji | Display Name          |
| ------------- | --------- | ----- | --------------------- |
| feature       | feat      | 🚀    | Feature               |
| fix           | bug       | 🐛    | Bugfix                |
| chore         | chore     | 🔧    | Chore                 |
| documentation | docs      | 📚    | Documentation         |
| refactor      | refactor  | ♻️    | Refactor              |
| test          | test      | 🧪    | Test                  |
| perf          | perf      | ⚡    | Performance           |
| ci            | ci        | 📦    | CI/CD                 |
| config        | cfg       | 🔧    | Configuration Change  |
| network       | net       | 🌐    | Network Change        |
| misc          | misc      | 📝    | Miscellaneous         |
| first-commit  | first     | 🏁    | First Commit          |
| milestone     | mile      | 🏆    | Milestone             |
| release       | release   | 🎯    | Release               |
| style         | style     | 🎨    | Style Change          |
| revert        | revert    | ⏪    | Revert                |
| merge         | merge     | 🔀    | Merge                 |
| security      | sec       | 🔒    | Security Fix          |
| build         | build     | 🏗️    | Build System          |
| deprecate     | depr      | 🗑️    | Deprecation           |
| ux            | ux        | 💡    | UX Improvement        |
| ui            | ui        | 🖼️    | UI Update             |
| hotfix        | hotfix    | 🚑    | Hotfix                |
| lint          | lint      | 🧹    | Linting               |
| env           | env       | 🌱    | Environment Setup     |
| legal         | legal     | 📄    | Legal/Compliance      |
| infra         | infra     | 🏭    | Infrastructure Change |
| i18n          | intl      | 🌍    | Internationalization  |
| analytics     | analytics | 📊    | Analytics/Tracking    |
| rollback      | rollback  | ↩️    | Rollback              |
| prototype     | proto     | 🧪    | Prototype/Experiment  |
| log           | log       | 📝    | Logging               |
| monitoring    | mon       | 📈    | Monitoring            |
| bump          | version   | 🔖    | Version Bump          |

🛠 Requirements
Go 1.16+

Git installed and available in PATH

$EDITOR environment variable set (default: vim)
