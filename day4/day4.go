package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4() (int, error) {
	f, err := os.Open("./day4/input.txt")
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("failed to scan file: %w", err)
	}

	var total int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if IsXMas(grid, i, j) {
				total++
			}
		}
	}
	return total, nil
}

func FindWordInDirection(grid [][]rune, x, y, dx, dy int, target string) bool {
	for i := 0; i < len(target); i++ {
		nextRow, nextCol := x+dx*i, y+dy*i

		outOfBounds := nextRow < 0 || nextRow >= len(grid) ||
			nextCol < 0 || nextCol >= len(grid[0])
		mismatch := !outOfBounds && grid[nextRow][nextCol] != rune(target[i])

		if outOfBounds || mismatch {
			return false
		}
	}
	return true
}

func IsXMas(grid [][]rune, row, col int) bool {
	if row-1 < 0 || row+1 >= len(grid) || col-1 < 0 || col+1 >= len(grid[0]) {
		return false
	}

	topLeftToBottomRight := (grid[row-1][col-1] == 'M' && grid[row][col] == 'A' && grid[row+1][col+1] == 'S') ||
		(grid[row-1][col-1] == 'S' && grid[row][col] == 'A' && grid[row+1][col+1] == 'M')

	topRightToBottomLeft := (grid[row-1][col+1] == 'M' && grid[row][col] == 'A' && grid[row+1][col-1] == 'S') ||
		(grid[row-1][col+1] == 'S' && grid[row][col] == 'A' && grid[row+1][col-1] == 'M')

	return topLeftToBottomRight && topRightToBottomLeft
}
