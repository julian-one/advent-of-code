package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() (int, error) {
	f, err := os.Open("./day2/input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var safe int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		level := strings.Fields(line)

		int_level := make([]int, len(level))
		for i, v := range level {
			int_level[i], _ = strconv.Atoi(v)
		}

		if IsSafeWithDampener(int_level) {
			safe++
		}
	}
	return safe, nil
}

func IsSafeWithDampener(level []int) bool {
	if IsSafe(level) {
		return true
	}

	for i := 0; i < len(level); i++ {
		modified := make([]int, 0, len(level)-1)
		modified = append(modified, level[:i]...)
		modified = append(modified, level[i+1:]...)

		if IsSafe(modified) {
			return true
		}
	}
	return false
}

func IsSafe(level []int) bool {
	direction := 0
	for i := 1; i < len(level); i++ {
		diff := level[i] - level[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		} else {
			if (direction == 1 && diff < 0) || (direction == -1 && diff > 0) {
				return false
			}
		}
	}
	return true
}
