package main

import (
	"fmt"

	"github.com/julian-one/advent-of-code/day1/part1"
)

func main() {
	answer, err := part1.Solution()
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 1 Part 1 Solution; sum of all of the calibration values:", answer)
}
