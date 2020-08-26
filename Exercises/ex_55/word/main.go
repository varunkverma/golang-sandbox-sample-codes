// Package word contains fucntions to count words in a quote in different ways
package word

import "strings"

// UseCount counts words and returns frequency for all worlds
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count simply returns of splitted words in a quote
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}