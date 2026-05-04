package task_1431

func findKidWithGreatestCandies(candies []int) int {
	if len(candies) < 1 {
		return 0
	}

	kidWithGreatestIdx := 0
	
	for i, c := range candies {
		if c > candies[kidWithGreatestIdx] {
			kidWithGreatestIdx = i
		}
	}

	return candies[kidWithGreatestIdx]
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
    greatestCandyCount := findKidWithGreatestCandies(candies)
	result := make([]bool, len(candies))

	for i, c := range candies {
		if (c + extraCandies) >= greatestCandyCount {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}