package utils

import "strings"

func ReplaceSpaces(input string) string {
	return strings.ReplaceAll(input, " ", "_")
}
