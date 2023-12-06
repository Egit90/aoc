package day1

import (
	"strconv"
	"unicode"

	"github.com/elie90/aoc/files"
)

var days = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   0,
}

func Day1() {

	file := files.NewFile("2023/day1/test")
	read := file.ReadFile()
	sum := 0

	for _, line := range read {
		tmp := extractWordsAndNumbers(line)
		dig, _ := strconv.Atoi(tmp)
		sum += dig
	}
	println(sum)
}

var tokens = map[string]string{
	"0": "zero",
	"1": "one",
	"2": "two",
	"3": "three",
	"4": "four",
	"5": "five",
	"6": "six",
	"7": "seven",
	"8": "eight",
	"9": "nine",
}

func checkIfDidgit(value rune) (bool, string) {
	if unicode.IsDigit(value) {
		return true, string(value)
	}
	return false, ""

}

func extractWordsAndNumbers(s string) string {
	res := []string{}

	for i, value := range s {

		if ok, digit := checkIfDidgit(rune(value)); ok {
			res = append(res, digit)
			continue
		}

		for key, mapValue := range tokens {

			if i+len(mapValue) > len(s) {
				continue
			}

			if s[i:i+len(mapValue)] == mapValue {
				res = append(res, key)
				continue
			}

		}

	}

	return res[0] + res[len(res)-1]

}
