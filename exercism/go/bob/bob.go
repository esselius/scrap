package bob

import (
	"strings"
	"unicode"
)

func silence(s string) bool {
	return s == ""
}

func nonsenseQuestion(s string) bool {
	return nonsense(s) && question(s)
}

func nonsense(s string) bool {
	for _, v := range s {
		if unicode.IsLetter(v) {
			return false
		}
	}

	return true
}

func yellingQuestion(s string) bool {
	return yelling(s) && question(s)
}

func yelling(s string) bool {
	for _, v := range s {
		if unicode.IsLower(v) {
			return false
		}
	}

	return true
}

func question(s string) bool {
	return strings.HasSuffix(s, "?")
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	switch {
	case silence(remark):
		return "Fine. Be that way!"
	case nonsenseQuestion(remark):
		return "Sure."
	case nonsense(remark):
		return "Whatever."
	case yellingQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case yelling(remark):
		return "Whoa, chill out!"
	case question(remark):
		return "Sure."
	}

	return "Whatever."
}
