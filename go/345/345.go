package task_345


func reverseVowels(s string) string {
	vowels := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	runes := []rune(s)
	idx1, idx2 := 0, len(runes) - 1
	
	for idx1 < idx2 {
		_, exists := vowels[runes[idx1]]
		if !exists {
			idx1++
		}
		_, exists2 := vowels[runes[idx2]]
		if !exists2 {
			idx2--
		}

		if exists && exists2 {
			runes[idx1], runes[idx2] = runes[idx2], runes[idx1]
			idx1++
			idx2--
		}
	}

	return string(runes)
}