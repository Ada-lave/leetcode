package task_3898

func findDegrees(matrix [][]int) []int {
	output := []int{}
	for _, row := range matrix {
		count := 0
		for _, value := range row {
			count += value
		}
		output = append(output, count)
	}

	return output
}