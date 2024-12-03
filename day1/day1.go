// https://adventofcode.com/2024/day/1

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func abs(number int) int {
	if number < 0 {
		return number * -1
	}
	return number
}

func main() {

	/**
	Approach:

	1. Read file line by line and add items on each line to 2 different arrays
	2. Sort the arrays (ascending order)
	3. Loop through the arrays, calculate the difference and sum them
	**/
	file, err := os.Open("day1/day1_test_case")
	check(err)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var firstList []int
	var secondList []int
	var currentLine []string
	var totalDistance int = 0

	for fileScanner.Scan() {
		currentLine = strings.Fields(fileScanner.Text())
		number, err := strconv.Atoi(currentLine[0])
		check(err)
		firstList = append(firstList, number)

		number2, err2 := strconv.Atoi(currentLine[1])
		check(err2)
		secondList = append(secondList, number2)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	for i := 0; i < len(firstList); i++ {
		totalDistance += abs(firstList[i] - secondList[i])
	}

	fmt.Println(totalDistance)

	file.Close()
}
