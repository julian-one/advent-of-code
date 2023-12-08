package main

import (
	"fmt"

	"github.com/julian-one/advent-of-code/day2"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func run() error {
	// answer, err := day1.Solution()
	answer, err := day2.Solution()
	if err != nil {
		panic(err)
	}
	// fmt.Println("Day 1 Solution; sum of all of the calibration values:", answer)
	fmt.Println("Day 2 Solution; sum of the IDs of valid games", answer)
	return nil
}
