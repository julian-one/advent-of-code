package day3

import (
	"os"
	"strconv"
)

func Day3() (int, error) {
	f, err := os.ReadFile("./day3/input.txt")
	if err != nil {
		return 0, err
	}
	input := string(f)

	enabled := true
	total := 0

	for i := 0; i < len(input); i++ {
		// do() and don't()
		if i+4 < len(input) && input[i:i+4] == "do()" {
			enabled = true
			i += 4
		} else if i+7 < len(input) && input[i:i+7] == "don't()" {
			enabled = false
			i += 7
		}

		// check for "mul("
		if i+4 < len(input) && input[i:i+4] == "mul(" {
			if !enabled {
				continue
			}
			i += 4

			// read the first number
			start := i
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			x, err := strconv.Atoi(input[start:i])
			if err != nil {
				continue
			}

			// check for a comma
			if i >= len(input) || input[i] != ',' {
				continue
			}
			i++

			// read the second number
			start = i
			for i < len(input) && input[i] >= '0' && input[i] <= '9' {
				i++
			}
			y, err := strconv.Atoi(input[start:i])
			if err != nil {
				continue
			}

			// check for closing perenthesis
			if i < len(input) && input[i] != ')' {
				continue
			}
			total += x * y
		}
	}
	return total, nil
}
