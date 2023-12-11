package main

import (
	"fmt"

	"github.com/julian-one/advent-of-code/day3"
)

func main() {
	err := run()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
}

func run() error {
	part1, err := day3.Solution()
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 3 Part 1 Solution", part1)
	return nil
}
