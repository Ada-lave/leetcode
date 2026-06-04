package task_443

import (
	"strconv"
)

func compress(chars []byte) int {
	if len(chars) == 1 {
		return 1
	}

	writePtr := 0
	readPtr := 0
	current := chars[writePtr]
	counter := 0

	for readPtr < len(chars) {
		if chars[readPtr] != current {
			chars[writePtr] = current
			writePtr++

			if counter > 1 {
				counterAtoi := strconv.Itoa(counter)
				for _, num := range counterAtoi {
					chars[writePtr] = byte(num)
					writePtr++

				}
			}
			counter = 1
			current = chars[readPtr]

		} else {
			counter++
		}

		readPtr++
	}

	chars[writePtr] = current
	writePtr++

	if counter > 1 {
		counterAtoi := strconv.Itoa(counter)
		for _, num := range counterAtoi {
			chars[writePtr] = byte(num)
			writePtr++

		}
	}

	return writePtr
}
