package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

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
	answer := part1()
	return answer, nil
}

func part1() int {
	lines, err := utils.ReadFile("day3/input.txt")
	if err != nil {
		log.Panic(err)
	}
	g, nm := create2DGrid(lines)
	neighbors := concatenateDigits(g, nm)
	fmt.Println("neighbors:", neighbors)
	answer := utils.SumArray(neighbors)
	return answer
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

func concatenateDigits(grid map[Point]rune, nm map[int][]NumberRange) []int {
	var concatenatedDigits []int

	for y, ranges := range nm {
		for _, r := range ranges {
			numberStr := ""
			for x := r.xStart; x <= r.xEnd; x++ {

				if char, ok := grid[Point{y, x}]; ok && unicode.IsDigit(char) {
					numberStr += string(char)
				}
			}
			if numberStr != "" {
				num, err := strconv.Atoi(numberStr)
				if err != nil {
					log.Panic(err)
				}
				fmt.Println("Number:", num, "at", Point{y, r.xStart})

				shouldAdd := false
				for x := r.xStart; x <= r.xEnd; x++ {
					ok := checkNeighbors(Point{y, x}, grid)
					if ok {
						shouldAdd = true
					}
				}
				if shouldAdd {
					concatenatedDigits = append(concatenatedDigits, num)
				}
			}
		}
	}

	return concatenatedDigits
}

func checkNeighbors(p Point, grid map[Point]rune) bool {

	neighbors := []Point{
		{x: p.x - 1, y: p.y},
		{x: p.x + 1, y: p.y},
		{x: p.x, y: p.y - 1},
		{x: p.x, y: p.y + 1},

		{x: p.x - 1, y: p.y - 1},
		{x: p.x + 1, y: p.y - 1},
		{x: p.x - 1, y: p.y + 1},
		{x: p.x + 1, y: p.y + 1},
	}

	for _, neighbor := range neighbors {
		if value, ok := grid[neighbor]; ok {
			if !unicode.IsLetter(value) && !unicode.IsDigit(value) && value != '.' {

				fmt.Println("Symbol at Neighbor:", neighbor, "Value:", string(value))
				return true
			}
		}
	}
	return false
}
