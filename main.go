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
	part2, err := day2.Solution()
	if err != nil {
		panic(err)
	}
	fmt.Println("Day 2 Part 2 Solution", part2)
	return nil
}
