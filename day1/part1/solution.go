package part1

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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
	var sum int
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading the string %s", err)
			break
		}
		nums := findAllIntsInLine(line)
		n, err := processIntList(nums)
		if err != nil {
			log.Panic(err)
		}
		calibrationValues = append(calibrationValues, n)
	}
	sum = sumArray(calibrationValues)
	return sum, nil
}

func readFile(path string) (*bufio.Reader, *os.File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, err
	}
	r := bufio.NewReader(f)
	return r, f, nil
}

func findAllIntsInLine(line string) []int {
	var nums []int
	for _, el := range line {
		// use ASCII value to check if it is a digit
		if el >= '0' && el <= '9' {
			// from ASCII value to int
			// 0 has ASCII value of 48
			// e.g. 3's ASCII value is 51, 51 - 48 = 3
			digit := int(el - '0')
			nums = append(nums, digit)
		}
	}
	return nums
}

func processIntList(list []int) (int, error) {
	if len(list) == 1 {
		c, err := combineInts(list[0], list[0])
		if err != nil {
			log.Panic(err)
		}
		return c, nil
	}

	a, b := list[0], list[len(list)-1]
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

func sumArray(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
