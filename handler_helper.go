package main

import "strings"

// Remove punctuation in title
func removePunc(s string) string {
	re := []string{}
	for _, char := range s {
		if (char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			(char == ' ') {
			re = append(re, string(char))
		}
	}
	return strings.Join(re, "")
}
