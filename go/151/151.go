package task_151

import "strings"

// "a good   example"
func reverseWords(s string) string {
    result := strings.Builder{}
	
	count := 0
	whiteSpace := byte(' ')
	for i := len(s) - 1; i >= 0; i-- {
		if count == 0 && s[i] == whiteSpace {
			continue
		} else if s[i] != whiteSpace {
			count++
		}
		if count > 0 && (s[i] == whiteSpace || i == 0) {
			result.WriteString(s[i + 1:i+count+1])
			count = 0
			if i != 0 {
				result.WriteByte(whiteSpace)
			}
		}
	}

	return strings.TrimSpace(result.String())
}