package main

import "strings"
import "unicode"

// containsEmoji returns true if the string contains characters commonly used for emojis.
// This is a conservative check covering major emoji blocks and special selectors.
func containsEmoji(s string) bool {
    for _, r := range s {
        switch {
        case r == 0x200D: // ZWJ
            return true
        case r == 0xFE0F: // VS16
            return true
        case r >= 0x1F300 && r <= 0x1F5FF: // Misc Symbols and Pictographs
            return true
        case r >= 0x1F600 && r <= 0x1F64F: // Emoticons
            return true
        case r >= 0x1F680 && r <= 0x1F6FF: // Transport and Map
            return true
        case r >= 0x1F900 && r <= 0x1F9FF: // Supplemental Symbols and Pictographs
            return true
        case r >= 0x1FA70 && r <= 0x1FAFF: // Symbols & Pictographs Extended-A
            return true
        case r >= 0x2600 && r <= 0x26FF: // Misc Symbols
            return true
        case r >= 0x2700 && r <= 0x27BF: // Dingbats
            return true
        case r >= 0x1F1E6 && r <= 0x1F1FF: // Regional Indicator (flags)
            return true
        default:
            // Heuristic: many emojis are 'So' (Symbol, other). Avoid false positives by limiting
            // to codepoints above Basic Multilingual Plane.
            if r > unicode.MaxASCII && unicode.In(r, unicode.So) {
                return true
            }
        }
    }
    return false
}

// removeCommentLines removes lines starting with '#' and trims whitespace. Skips empty lines.
func removeCommentLines(content string) string {
	var result []string
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		result = append(result, trimmed)
	}

	return strings.Join(result, "\n")
}
