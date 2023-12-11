package day2

import (
	"log"
	"strconv"
	"strings"

	"github.com/julian-one/advent-of-code/utils"
)

const (
	red   = 12
	green = 13
	blue  = 14
)

type Game struct {
	id   int
	cube []Cube
}

type Cube struct {
	color string
	count int
}

func Solution() (int, error) {
	powerOfGames := calculate()
	part2 := utils.SumArray(powerOfGames)
	return part2, nil
}

func calculate() []int {
	lines, err := utils.ReadFile("day2/input.txt")
	if err != nil {
		log.Panic(err)
	}

	var gamePowers []int

	for _, line := range lines {
		_, maxCounts := processGame(line)

		gamePower := powerCubeSet(maxCounts)
		gamePowers = append(gamePowers, gamePower)

	}
	return gamePowers
}

func processGame(game string) (int, map[string]int) {

	parts := strings.Split(game, ":")
	if len(parts) != 2 {
		log.Panic("invalid game format")
	}

	gameIdPart := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))
	gameId, _ := strconv.Atoi(gameIdPart)

	cubeSets := strings.Split(parts[1], ";")
	maxCounts := make(map[string]int)

	for _, part := range cubeSets {

		cubeColors := strings.Split(strings.TrimSpace(part), ",")

		for _, cubeColor := range cubeColors {

			colorData := strings.Fields(strings.TrimSpace(cubeColor))

			if len(colorData) == 2 {
				count, _ := strconv.Atoi(colorData[0])
				color := colorData[1]

				currentMin, exists := maxCounts[color]
				if !exists || count > currentMin {
					maxCounts[color] = count
				}
			}
		}
	}
	return gameId, maxCounts
}

func (c Cube) isValidGame() bool {
	switch c.color {
	case "red":
		return c.count <= red
	case "green":
		return c.count <= green
	case "blue":
		return c.count <= blue
	}
	return false
}

func powerCubeSet(cubeSet map[string]int) int {
	power := 1
	for _, v := range cubeSet {
		power *= v
	}
	return power
}
