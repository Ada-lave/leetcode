package task_151

import (
	"slices"
	"strings"
)

// "a good   example"
func reverseWords(s string) string {
    result := []string{}
	
	i := 0
	currentString := strings.Builder{}
	for i < len(s) {
		if s[i] != ' ' {
			currentString.WriteByte(s[i])
		} else {
			if currentString.Len() > 0 {
				result = append(result, currentString.String())
			}
			currentString.Reset()
		}
		i++
	}
	if currentString.Len() > 0 {
		result = append(result, currentString.String())
	}
	slices.Reverse(result)
	return strings.Join(result, " ")
}