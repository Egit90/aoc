package day3

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/elie90/aoc/files"
	"github.com/elie90/aoc/types"
)

type game struct {
	puzzle        []string
	listOfMatches map[int][][]int
	startsIndex   map[int][][]int
	res           int
}
type cords struct {
	lineIndex int
	charIndex int
}

func Day3(partNumber types.PartNumber) {
	file := files.NewFile("2023/day3/input")
	read := file.ReadFile()

	if types.One == partNumber {
		p1 := &game{puzzle: read, listOfMatches: make(map[int][][]int)}
		p1.getListOfNumbersAndIndexes()
		p1.analyzeMatchesWithinLine()
		fmt.Println(p1.res)
		return
	}

	//part 2

	p1 := &game{puzzle: read, listOfMatches: make(map[int][][]int)}
	p1.getListOfNumbersAndIndexes()
	p1.getListOfStarsIndexs()

}

func (p *game) getListOfNumbersAndIndexes() {
	for lineIndex, line := range p.puzzle {
		re := regexp.MustCompile(`[\d]+`)
		matches := re.FindAllIndex([]byte(line), -1)
		p.listOfMatches[lineIndex] = matches

	}
}

func (p *game) getListOfStarsIndexs() {
	for lineIndex, line := range p.puzzle {
		re := regexp.MustCompile(`\*`)
		matches := re.FindAllIndex([]byte(line), -1)
		p.startsIndex[lineIndex] = matches

	}
}

func (p *game) analyzeMatchesWithinLine() {

	for lineIndex, matches := range p.listOfMatches {
		for _, match := range matches {
			startIdx, endIdx := match[0], match[1]
			p.addToResult(lineIndex, startIdx, endIdx)
		}
	}

}

func (p *game) addToResult(lineIndex, startIdx, endIdx int) {
	if p.checkIfShouldBeAdded(lineIndex, startIdx, endIdx-1) {
		substring := p.puzzle[lineIndex][startIdx:endIdx]
		if number, err := strconv.Atoi(substring); err == nil {
			p.res += number
		}

	}
}

func (p *game) checkIfShouldBeAdded(lineIndex, leftIndex, rightIndex int) bool {

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
		if Char, err := p.getValueAtIndex(v.lineIndex, v.charIndex); err == nil {
			if isDot(Char, 0) {
				continue
			}
			return true
		}

	}

	return false
}

func (p *game) lineExists(line int) ([][]int, bool) {
	if line > 0 && line < len(p.listOfMatches)-1 {
		return p.listOfMatches[line], true
	}
	return nil, false
}

func (p *game) getValueAtIndex(lineIndex, index int) (string, error) {

	if lineIndex < 0 || lineIndex >= len(p.puzzle) {
		return "", errors.New("out of bound")
	}

	s := (p.puzzle)[lineIndex]

	if index < 0 || index >= len(s) {
		return "", errors.New("out of bound")
	}
	return string(s[index]), nil
}

func isDot(s string, index int) bool {
	if index < len(s) {
		return string(s[index]) == "."
	}

	return true
}
