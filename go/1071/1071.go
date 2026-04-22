package task_1071

import "strings"

func FindAllSubStrings(text string) []string {
	subStringMap := []string{}
	for i := 0; i < len(text) - 1; i++ {
		for j := i + 1; j < len(text); j++ {
			subStringMap = append(subStringMap, text[i:j + 1])
		
		}
	}

	if len(text) == 1 {
		subStringMap = append(subStringMap, text)
	}

	return subStringMap
}

func FindSubStrCount(text string, sub string) int {
	count := 1
	for i := 0; i < len(text) - 1; i++ {
		for j := i + 1; j < len(text); j++ {
			if text[i:j + 1] == sub {
				count++
			}
		
		}
	}
	return count
}

func gcdOfStrings(str1 string, str2 string) string {
    
	subMap1 := FindAllSubStrings(str2)

	for _, v := range subMap1 {
		subCount := FindSubStrCount(str1, v)
		if strings.Repeat(v, subCount) == str1 {
			return v
		}
	}

	return ""
}