package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elie90/aoc/files"
)

func createDefaultValuesMap() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func Day2(isPartTwo bool) {
	file := files.NewFile("2023/day2/input")
	read := file.ReadFile()
	total := 0

	var processFunction processFunc
	if isPartTwo {
		processFunction = partTwo
	} else {
		processFunction = partOne
	}

	for _, line := range read {
		total += processLine(line, processFunction)
	}

	fmt.Println(total)
}

type processFunc func([]string, int) int

func processLine(s string, processFunction processFunc) int {
	input := strings.Split(s, ":")
	gameNumber, _ := strconv.Atoi(strings.Split(input[0], " ")[1])
	gameInput := strings.Split(input[1], ";")

	return processFunction(gameInput, gameNumber)
}

func partOne(gameInput []string, gameNumber int) int {
	for _, v := range gameInput {
		fi := strings.Split(v, ",")
		for _, line := range fi {
			thisNumber, thePick := separateBySpace(line)
			myMap := createDefaultValuesMap()
			myMap[thePick] -= thisNumber
			if myMap[thePick] < 0 {
				return 0
			}
		}
	}
	return gameNumber
}

func partTwo(gameInput []string, gameNumber int) int {
	myMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	for _, v := range gameInput {
		fi := strings.Split(v, ",")
		for _, line := range fi {
			thisNumber, thePick := separateBySpace(line)
			if myMap[thePick] < thisNumber {
				myMap[thePick] = thisNumber
			}
		}
	}

	total := 1
	for _, v := range myMap {
		total *= v
	}
	return total
}

func separateBySpace(s string) (int, string) {
	newLine := strings.TrimSpace(s)
	line := strings.Split(newLine, " ")

	_tmpNum := line[0]
	num, err := strconv.Atoi(_tmpNum)
	if err != nil {
		panic(fmt.Sprintf("failed to convert string to int %s", s))
	}

	return num, line[1]
}
