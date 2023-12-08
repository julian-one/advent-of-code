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

func Solution() error {
	processFileByLine()
	return nil
}

func readFile(path string) (*bufio.Scanner, *os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	s := bufio.NewScanner(f)
	return s, f, nil
}

func processFileByLine() {
	s, f, err := readFile("day2/input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	for s.Scan() {
		line := s.Text()
		m := makeGameMap(line)
		fmt.Println(m)
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}

func makeGameMap(game string) map[int]map[string]int {
	gameData := make(map[int]map[string]int)

	parts := strings.Split(game, ":")

	if len(parts) != 2 {
		log.Panic("invalid game format")
	}

	gameIdPart := strings.TrimSpace(strings.TrimPrefix(parts[0], "Game"))
	gameId, _ := strconv.Atoi(gameIdPart)

	gameData[gameId] = make(map[string]int)
	colorParts := strings.Split(parts[1], ";")

	for _, part := range colorParts {
		items := strings.Split(strings.TrimSpace(part), ",")

		for _, item := range items {
			colorData := strings.Fields(strings.TrimSpace(item))

			if len(colorData) == 2 {
				count, _ := strconv.Atoi(colorData[0])
				color := colorData[1]
				gameData[gameId][color] += count
			}
		}
	}

	return gameData
}
