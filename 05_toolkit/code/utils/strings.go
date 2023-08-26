package utils

import "strings"

// MakeExcited transforms a sentence into all caps with an exclamation mark
func MakeExcited(txt string) string {
	return strings.ToUpper((txt)) + "!"
}
