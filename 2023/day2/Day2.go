package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/elie90/aoc/files"
)

func createDefValuesMap() map[string]int {
	return map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
}

func Day2() {

	file := files.NewFile("2023/day2/input")
	read := file.ReadFile()
	total := 0
	for _, line := range read {
		if ok, game := processLine(line); ok {
			total += game
		}
	}
	fmt.Println(total)
}

func processLine(s string) (bool, int) {

	input := strings.Split(s, ":")

	gameNumber, _ := strconv.Atoi(strings.Split(input[0], " ")[1])

	// 3 blue, 4 red
	// 1 red, 2 green, 6 blue
	// 2 green
	gameInput := strings.Split(input[1], ";")

	for _, v := range gameInput {

		// 3 blue
		// 4 red
		fi := strings.Split(v, ",")

		for _, line := range fi {
			thisNumber, thePick := separateBySpace(line)
			myMap := createDefValuesMap()
			myMap[thePick] -= thisNumber
			if myMap[thePick] < 0 {
				return false, 0
			}

		}

	}

	return true, gameNumber
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
