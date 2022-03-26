package rules

import "strings"

// Unique returns true if all elements in the string are unique.
func Unique[T comparable](data []T) bool {
	uniqueValues := map[T]bool{}

	for _, value := range data {
		uniqueValues[value] = true
	}

	return len(data) == len(uniqueValues)
}

// NotBlank returns true if no element in the string is blank.
func NotBlank(data []string) bool {
	for _, genre := range data {
		if strings.TrimSpace(genre) == "" {
			return false
		}
	}
	return true
}
