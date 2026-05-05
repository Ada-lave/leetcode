package task_605

func canPlaceFlowers(flowerbed []int, n int) bool {
    for i := 0; i < len(flowerbed); i++ {
		if (i == 0 || flowerbed[i-1] != 1) && (flowerbed[i] == 0) && (i == len(flowerbed) -1 || flowerbed[i + 1] != 1) {
			n--
			i++
		}
	}

	return n <= 0
}