package answer

import (
	"fmt"
	"strings"
)

func shouldEscape(b byte) bool {
	switch b {
	case '_':
		return true
	case '*':
		return true
	case '[':
		return true
	case ']':
		return true
	case '(':
		return true
	case ')':
		return true
	case '~':
		return true
	case '`':
		return true
	case '>':
		return true
	case '#':
		return true
	case '+':
		return true
	case '-':
		return true
	case '=':
		return true
	case '|':
		return true
	case '{':
		return true
	case '}':
		return true
	case '.':
		return true
	case '!':
		return true
	}
	return false
}

func escape(str string) string {
	original := []byte(str)
	escaped := []byte{}
	prevPercent := false
	open := false
	for _, b := range original {
		if b == '%' {
			prevPercent = true
		}
		if prevPercent && b == '[' {
			open = true
		}
		if !open && shouldEscape(b) {
			escaped = append(escaped, '\\')
		}
		if open && b == ']' {
			open = false
		}
		escaped = append(escaped, b)
		if b != '%' {
			prevPercent = false
		}
	}
	return string(escaped)
}

func InsertMarkdown(str string) string {
	// Escape all characters
	str = escape(str)
	// Add markdown syntax
	str = fmt.Sprintf(str, "*", "_", "__", "~")
	// Remove possible error of fmt.Sprintf
	const extraError = "%!(EXTRA string=*, string=_, string=__, string=~)"
	str = strings.TrimSuffix(str, extraError)
	return str
}
