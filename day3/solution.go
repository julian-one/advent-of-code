package day3

import (
	"fmt"
	"log"
	"regexp"

	"github.com/julian-one/advent-of-code/utils"
)

type Point struct {
	y int
	x int
}

type NumberRange struct {
	xStart int
	xEnd   int
}

func Solution() (int, error) {
	part1()
	return 0, nil
}

func part1() {
	lines, err := utils.ReadFile("day3/input.txt")
	if err != nil {
		log.Panic(err)
	}
	g, nm := create2DGrid(lines)
	fmt.Println("grid:", g)
	fmt.Println("number ranges:", nm)
	checkNeighbors(g, nm)
}

func create2DGrid(lines []string) (map[Point]rune, map[int][]NumberRange) {

	twoDimensionalCharMap := make(map[Point]rune)
	numberRangeMap := make(map[int][]NumberRange)

	re := regexp.MustCompile(`\d+`)

	for y, line := range lines {
		for x, char := range line {
			twoDimensionalCharMap[Point{y, x}] = char
		}

		numbers := re.FindAllStringIndex(line, -1)
		for _, pair := range numbers {
			startIndex, endIndex := pair[0], pair[1]-1 // endIndex is adjusted to be inclusive
			numberRangeMap[y] = append(numberRangeMap[y], NumberRange{xStart: startIndex, xEnd: endIndex})
		}
	}

	return twoDimensionalCharMap, numberRangeMap
}

func checkNeighbors(grid map[Point]rune, nm map[int][]NumberRange) {
	for y, ranges := range nm {
		for _, r := range ranges {
			for x := r.xStart; x <= r.xEnd; x++ {
				fmt.Println("y:", y, "x:", x)
			}
		}
	}
}
