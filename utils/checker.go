package utils

import "unicode"

func IsInt(s string) bool {
    for i, c := range s {
        if (i == 0 && c == '-') {
            continue;
        }
        
        if !unicode.IsDigit(c) {
            return false
        }
    }
    return true
}