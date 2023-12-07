package part1

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func Solution() {
	readFileByLine()

}
func readFileByLine() error {
	f, err := os.Open("./day1/part1/input.txt")
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
		fmt.Print(line)
	}
	return nil
}

func findNumbersInLine(line string) int {
	return 10
}
