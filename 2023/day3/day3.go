package day3

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elie90/aoc/files"
	"github.com/elie90/aoc/types"
)

func Day3(partNumber types.PartNumber) {
	file := files.NewFile("2023/day3/input")
	read := file.ReadFile()

	if types.One == partNumber {
		play(read)
	}

}

func play(puzzle []string) {
	res := 0
	for lineIndex, line := range puzzle {

		re := regexp.MustCompile(`[\d]+`)
		matches := re.FindAllIndex([]byte(line), -1)

		for _, match := range matches {
			startIdx, endIdx := match[0], match[1]
			if ok := checkIfShouldBeAdded(puzzle, lineIndex, startIdx, endIdx-1); ok {
				substring := line[startIdx:endIdx]
				fmt.Printf("Add Number : %s \n", substring)
				if number, err := strconv.Atoi(substring); err == nil {
					res += number
				}

			}
		}

	}

	fmt.Println(res)
}

func isDot(s string, index int) bool {
	if index < len(s) {
		return string(s[index]) == "."
	}

	return true
}

type cords struct {
	lineIndex int
	charIndex int
}

func checkIfShouldBeAdded(puzzle []string, lineIndex int, leftIndex int, rightIndex int) bool {

	directions := map[string]cords{
		"leftSide":    {lineIndex, leftIndex - 1},
		"rightSide":   {lineIndex, rightIndex + 1},
		"upRight":     {lineIndex - 1, rightIndex},
		"upRight+1":   {lineIndex - 1, rightIndex + 1},
		"upLeft":      {lineIndex - 1, leftIndex},
		"upLeft+1":    {lineIndex - 1, leftIndex - 1},
		"DownRight":   {lineIndex + 1, rightIndex},
		"DownRight+1": {lineIndex + 1, rightIndex + 1},
		"DownLeft":    {lineIndex + 1, leftIndex},
		"DownLeft+1":  {lineIndex + 1, leftIndex - 1},
	}

	for i := leftIndex + 1; i < rightIndex; i++ {
		directions[fmt.Sprintf("middleUp%d", i)] = cords{lineIndex - 1, i}
		directions[fmt.Sprintf("middledown%d", i)] = cords{lineIndex + 1, i}
	}

	for _, v := range directions {
		if Char, err := getValueAtIndex(puzzle, v.lineIndex, v.charIndex); err == nil {
			if isDot(Char, 0) {
				continue
			}
			return true
		}

	}

	return false
}

func getValueAtIndex(puzzle []string, lineIndex int, index int) (string, error) {

	if lineIndex < 0 || lineIndex >= len(puzzle) {
		return "", errors.New("out of bound")
	}

	s := puzzle[lineIndex]

	if index < 0 || index >= len(s) {
		return "", errors.New("out of bound")
	}
	return string(s[index]), nil
}
