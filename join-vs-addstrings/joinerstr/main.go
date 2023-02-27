// Package joinerstr for all string manipulations
package joinerstr

import "strings"

// AddStr returns the sentence
func AddStr(s []string) string {
	k := s[0]
	for _, v := range s[1:] {
		k += " "
		k += v
	}
	return k
}

// Joint returns the sentence
func Joint(s []string) string {
	sf := strings.Join(s, " ")
	return sf
}
