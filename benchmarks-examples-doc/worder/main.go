// Utility functions for all the word counts
package worder

import "strings"

// UseCount returns number of times each word is used
func UseCount(s string) map[string]int {
	sf := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range sf {
		m[v]++
	}
	return m
}

// Counter returns number of words used
func Counter(s string) int {
	sf := strings.Fields(s)
	return len(sf)
}
