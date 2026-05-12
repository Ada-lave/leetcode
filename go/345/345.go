package task_345

func reverseVowels(s string) string {
	vowels := map[string]struct{}{
		"a": {},
		"e": {},
		"i": {},
		"o": {},
		"u": {},
		"A": {},
		"E": {},
		"I": {},
		"O": {},
		"U": {},
	}

	idx1, idx2 := 0, len(s)

	for idx1 < len(s) && idx2 != 0 {
		if _, exists := vowels[string(s[idx1])]; exists {
			
		}
	}

	return s
}