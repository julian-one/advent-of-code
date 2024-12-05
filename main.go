package main

import (
	"advent-of-code-2024/day1"
	"advent-of-code-2024/day2"
	"advent-of-code-2024/day3"
	"advent-of-code-2024/day4"
	"log"
)

func main() {
	// Day 1
	l, r, err := day1.ParseInput()
	if err != nil {
		log.Fatalf("failed to parse input for day 1: %v", err)
	}
	d1 := day1.Day1(l, r)
	if err != nil {
		log.Fatalf("failed to solve day 1: %v", err)
	}
	log.Printf("day 1 answer: %d", d1)
	// Day 2
	d2, err := day2.Day2()
	if err != nil {
		log.Fatalf("failed to solve day 2: %v", err)
	}
	log.Printf("day 2 answer: %d", d2)
	// Day 3
	d3, err := day3.Day3()
	if err != nil {
		log.Fatalf("failed to solve day 3: %v", err)
	}
	log.Printf("day 3 answer: %d", d3)
	d4, err := day4.Day4()
	if err != nil {
		log.Fatalf("failed to solve day 4: %v", err)
	}
	log.Printf("day 4 answer: %d", d4)
}
