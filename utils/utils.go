package utils

import (
	"os"
	"regexp"
)

// ========================
// FILES IO
// ========================
func ReadFile(title string) string {
	data, err := os.ReadFile(title)
	if err != nil {
		panic("Error reading file")
	}
	return string(data)
}

// ========================
// STRINGS
// ========================
func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}
