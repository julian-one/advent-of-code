package main

import (
	"fmt"

	"github.com/julian-one/advent-of-code/day1"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func run() error {
	answer, err := day1.Solution()
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 1 Solution; sum of all of the calibration values:", answer)
	return nil
}
