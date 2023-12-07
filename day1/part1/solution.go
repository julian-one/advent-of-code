package part1

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solution() error {
	err := readFileByLine()
	if err != nil {
		return err
	}
	return nil
}

func readFileByLine() error {
	f, err := os.Open("./day1/part1/input-trunc.txt")
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading the string %s", err)
			break
		}
		fmt.Printf("readFileByLine, line: %s", line)
		findNumbersInLine(line)
	}
	return nil
}

func findNumbersInLine(line string) (int, int, error) {

	for _, el := range line {
		// use ASCII to check if it is a digit
		if el >= '0' && el <= '9' {
			fmt.Println("findNumbersInLine, el:", el)
		}
	}

	return 10, 13, nil
}
