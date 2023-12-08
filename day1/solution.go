package day1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/julian-one/advent-of-code/utils"
)

func Solution() (int, error) {
	sum, err := processFileByLine()
	if err != nil {
		return -1, err
	}
	return sum, nil
}

func processFileByLine() (int, error) {
	r, f, err := readFile("./day1/input.txt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	var calibrationValues []int
	var calibrationSum int
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading the string %s", err)
			break
		}
		m1 := findSpelledOutDigits(line)
		m2 := findDigits(line, m1)
		smallest, largest := processMap(m2)
		calibrationValue, err := processInts(smallest, largest)
		if err != nil {
			log.Panic(err)
		}
		calibrationValues = append(calibrationValues, calibrationValue)
	}
	calibrationSum = utils.SumArray(calibrationValues)
	return calibrationSum, nil
}

func readFile(path string) (*bufio.Reader, *os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	r := bufio.NewReader(f)
	return r, f, nil
}

func findDigits(line string, m map[int]int) map[int]int {
	for i, el := range line {
		// use ASCII value to check if it is a digit
		if el >= '0' && el <= '9' {
			// from ASCII value to int
			// 0 has ASCII value of 48
			// e.g. 3's ASCII value is 51, 51 - 48 = 3
			digit := int(el - '0')
			m[i] = digit
		}
	}
	return m
}

func findSpelledOutDigits(line string) map[int]int {
	m := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	line = strings.ToLower(line)
	digitMap := make(map[int]int)

	for word, digit := range m {
		index := 0
		for {
			position := strings.Index(line[index:], word)
			if position == -1 {
				break
			}
			originalIndex := index + position

			digitAsNumber, _ := strconv.Atoi(digit)
			digitMap[originalIndex] = digitAsNumber
			index += position + len(word)
		}
	}
	return digitMap
}

func processMap(m map[int]int) (int, int) {
	if len(m) == 0 {
		log.Panic("empty map?!")
	}

	var smallest, largest int
	first := true

	for key := range m {
		if first {
			smallest, largest = key, key
			first = false
		} else {
			if key < smallest {
				smallest = key
			}
			if key > largest {
				largest = key
			}
		}
	}
	return m[smallest], m[largest]
}

func processInts(a int, b int) (int, error) {
	c, err := combineInts(a, b)
	if err != nil {
		log.Panic(err)
	}
	return c, nil
}

func combineInts(a int, b int) (int, error) {
	str1 := strconv.Itoa(a)
	str2 := strconv.Itoa(b)
	combined := str1 + str2

	result, err := strconv.Atoi(combined)
	if err != nil {
		return -1, err
	}
	return result, nil
}
