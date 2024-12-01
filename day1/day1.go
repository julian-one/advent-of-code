package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ParseInput() ([]int64, []int64, error) {
	f, err := os.Open("./day1/input.txt")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open input file: %v", err)
	}
	defer f.Close()

	var left, right []int64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		left_num, _ := strconv.ParseInt(parts[0], 10, 64)
		right_num, _ := strconv.ParseInt(parts[1], 10, 64)

		left = append(left, left_num)
		right = append(right, right_num)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to read input file: %v", err)
	}

	return left, right, nil
}

func Day1(left, right []int64) (int64, error) {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	total := int64(0)
	for i := range left {
		distance := left[i] - right[i]
		if distance < 0 {
			distance = -distance
		}
		total += distance
	}
	return total, nil
}
