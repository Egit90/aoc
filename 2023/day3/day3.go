package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/elie90/aoc/files"
	"github.com/elie90/aoc/types"
)

type game struct {
	puzzle        []string
	listOfMatches map[int][]cords
	symboleIndex  map[int][]int
	res           int
}
type cords struct {
	startIdx int
	lastIdx  int
	Number   int
}

func Day3(partNumber types.PartNumber) {
	file := files.NewFile("2023/day3/input")
	read := file.ReadFile()

	if types.One == partNumber {
		p1 := &game{puzzle: read, listOfMatches: make(map[int][]cords), symboleIndex: make(map[int][]int)}
		p1.getListOfNumbersAndIndexes()
		p1.getListOfSymboleIndexs()
		p1.compare()
		fmt.Println(p1.res)
		return
	}
	p1 := &game{puzzle: read, listOfMatches: make(map[int][]cords), symboleIndex: make(map[int][]int)}
	p1.getListOfNumbersAndIndexes()
	p1.getListOfStarIndex()
	p1.compareForPart2()
	fmt.Println(p1.res)
}

func (p *game) getListOfNumbersAndIndexes() {
	for lineIndex, line := range p.puzzle {
		re := regexp.MustCompile(`[\d]+`)
		matches := re.FindAllIndex([]byte(line), -1)
		for _, match := range matches {
			num, _ := strconv.Atoi(line[match[0]:match[1]])
			_tmp := cords{
				startIdx: match[0],
				lastIdx:  match[1] - 1,
				Number:   num,
			}
			p.listOfMatches[lineIndex] = append(p.listOfMatches[lineIndex], _tmp)
		}
	}
}

func (p *game) getListOfSymboleIndexs() {
	re := regexp.MustCompile(`[^0-9.]`)
	for lineIndex, line := range p.puzzle {
		matches := re.FindAllIndex([]byte(line), -1)
		for _, match := range matches {
			p.symboleIndex[lineIndex] = append(p.symboleIndex[lineIndex], match[0])
		}
	}
}

func (p *game) getListOfStarIndex() {
	re := regexp.MustCompile(`[\*]`)
	for lineIndex, line := range p.puzzle {
		matches := re.FindAllIndex([]byte(line), -1)
		for _, match := range matches {
			p.symboleIndex[lineIndex] = append(p.symboleIndex[lineIndex], match[0])
		}
	}
}

func (p *game) compare() {
	for symbolY, symbolXList := range p.symboleIndex {
		for _, symbolX := range symbolXList {
			p.checkForNumbers(symbolX, symbolY)
		}
	}
}

func (p *game) checkForNumbers(symboleX, symboleY int) {

	li := [][]cords{}

	if line, ok := p.listOfMatches[symboleY]; ok {
		li = append(li, line)
	}
	if lineAbove, ok := p.listOfMatches[symboleY-1]; ok {
		li = append(li, lineAbove)
	}
	if lineBellow, ok := p.listOfMatches[symboleY+1]; ok {
		li = append(li, lineBellow)
	}

	for _, cords := range li {
		for _, cord := range cords {
			if symboleX == cord.startIdx || symboleX == cord.startIdx-1 || symboleX == cord.startIdx+1 ||
				symboleX == cord.lastIdx || symboleX == cord.lastIdx+1 || symboleX == cord.lastIdx-1 {
				p.res += cord.Number
			}
		}
	}
}

func (p *game) compareForPart2() {
	for symbolY, symbolXList := range p.symboleIndex {
		for _, symbolX := range symbolXList {
			p.checkForNumbersPart2(symbolX, symbolY)
		}
	}
}

func (p *game) checkForNumbersPart2(symboleX, symboleY int) {

	li := [][]cords{}

	if line, ok := p.listOfMatches[symboleY]; ok {
		li = append(li, line)
	}
	if lineAbove, ok := p.listOfMatches[symboleY-1]; ok {
		li = append(li, lineAbove)
	}
	if lineBellow, ok := p.listOfMatches[symboleY+1]; ok {
		li = append(li, lineBellow)
	}

	myLi := []int{}
	for _, cords := range li {
		for _, cord := range cords {
			if symboleX == cord.startIdx || symboleX == cord.startIdx-1 || symboleX == cord.startIdx+1 ||
				symboleX == cord.lastIdx || symboleX == cord.lastIdx+1 || symboleX == cord.lastIdx-1 {
				myLi = append(myLi, cord.Number)
			}
		}
		if len(myLi) == 2 {
			p.res += (myLi[0] * myLi[1])
			myLi = []int{}
		}
	}

}
