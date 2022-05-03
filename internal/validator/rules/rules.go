package rules

import (
	"regexp"
	"strings"
	"unicode"
)

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

// In returns true if the data is found in the target list.
func In[T comparable](data T, target []T) bool {
	for _, element := range target {
		if element == data {
			return true
		}
	}
	return false
}

// MatchesPattern returns true if a string value matches a specific regexp pattern.
func MatchesPattern(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// NoWhiteSpace returns true if the string value contains no whitespace character.
func NoWhiteSpace(value string) bool {
	for _, char := range value {
		if unicode.IsSpace(char) {
			return false
		}
	}
	return true
}
