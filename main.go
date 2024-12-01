package main

import (
	"advent-of-code-2024/day1"
	"log"
)

func main() {
	l, r, err := day1.ParseInput()
	if err != nil {
		log.Fatalf("failed to parse input for day 1: %v", err)
	}
	answer := day1.Day1(l, r)
	if err != nil {
		log.Fatalf("failed to solve day 1: %v", err)
	}
	log.Printf("day 1 answer: %d", answer)
}
