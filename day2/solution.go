package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

func Solution() (int, error) {
	ids := processFileByLine()
	answer := sumArray(ids)
	return answer, nil
}

func readFile(path string) (*bufio.Scanner, *os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	s := bufio.NewScanner(f)
	return s, f, nil
}

func processFileByLine() []int {
	s, f, err := readFile("day2/input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var validGameIds []int

	for s.Scan() {
		line := s.Text()
		id, isValid := processGame(line)
		if isValid {
			validGameIds = append(validGameIds, id)
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	return validGameIds
}

func processGame(game string) (int, bool) {

	parts := strings.Split(game, ":")
	if len(parts) != 2 {
		log.Panic("invalid game format")
	}

	gameIdPart := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))
	gameId, _ := strconv.Atoi(gameIdPart)

	cubeSets := strings.Split(parts[1], ";")

	for _, part := range cubeSets {

		cubeColors := strings.Split(strings.TrimSpace(part), ",")

		for _, cubeColor := range cubeColors {

			colorData := strings.Fields(strings.TrimSpace(cubeColor))

			if len(colorData) == 2 {
				count, _ := strconv.Atoi(colorData[0])
				color := colorData[1]

				isValid := isValidGame(color, count)
				if !isValid {
					return gameId, false
				}
			}
		}
	}
	return gameId, true
}

func isValidGame(cubeColor string, cubeCount int) bool {
	switch cubeColor {
	case "red":
		return cubeCount <= red
	case "green":
		return cubeCount <= green
	case "blue":
		return cubeCount <= blue
	}
	return false
}

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
