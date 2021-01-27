package utils

import "unicode"

// IsInt is a function to check if a given string is "number"
func IsInt(s string) bool {
	if s == "" {
		return false
	}

	for i, c := range s {
		if i == 0 && c == '-' {
			continue
		}

		if !unicode.IsDigit(c) {
			return false
		}
	}

	return true
}
