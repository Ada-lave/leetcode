package task_1768

import "strings"

func mergeAlternately(word1 string, word2 string) string {
    resultString := strings.Builder{}
	i,j := 0, 0

	for i < len(word1) || j < len(word2) {
		if i < len(word1) {
			resultString.WriteByte(word1[i])
			i++
		}

		if j < len(word2) {
			resultString.WriteByte(word2[j])
			j++
		}
	}
	return resultString.String()
}