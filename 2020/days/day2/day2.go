package day2

import (
	"aoc2020/libs"
	"fmt"
	"strconv"
	"strings"
)

func p1IsValid(value string, char rune, min int, max int) bool {
	chars := []rune(value)
	count := 0
	for i := 0; i < len(chars); i++ {
		if chars[i] == char {
			count++
		}

		if count > max {
			return false
		}
	}

	return count >= min
}

func part1() {
	input := libs.ReadTxtFileLines("days/day2/input.txt")

	validCount := 0
	for i := 0; i < len(input); i++ {
		parts := strings.Split(strings.TrimSpace(input[i]), ":")
		if len(parts) != 2 {
			panic("this should be 2")
		}

		password := strings.TrimSpace(parts[1])
		parts = strings.Split(strings.TrimSpace(parts[0]), " ")

		if len(parts) != 2 {
			panic("this should also be 2")
		}

		character := strings.TrimSpace(parts[1])
		parts = strings.Split(strings.TrimSpace(parts[0]), "-")

		if len(parts) != 2 {
			panic("apparently this should also be 2")
		}
		val, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic("could not parse min value")
		}
		min := val
		val, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic("could not parse max value")
		}
		max := val

		if p1IsValid(password, []rune(character)[0], min, max) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func p2IsValid(value string, char rune, firstIndex int, secondIndex int) bool {
	chars := []rune(value)

	return (chars[firstIndex-1] == char && chars[secondIndex-1] != char) || (chars[firstIndex-1] != char && chars[secondIndex-1] == char)
}

func part2() {
	input := libs.ReadTxtFileLines("days/day2/input.txt")

	validCount := 0
	for i := 0; i < len(input); i++ {
		parts := strings.Split(strings.TrimSpace(input[i]), ":")
		if len(parts) != 2 {
			panic("this should be 2")
		}

		password := strings.TrimSpace(parts[1])
		parts = strings.Split(strings.TrimSpace(parts[0]), " ")

		if len(parts) != 2 {
			panic("this should also be 2")
		}

		character := strings.TrimSpace(parts[1])
		parts = strings.Split(strings.TrimSpace(parts[0]), "-")

		if len(parts) != 2 {
			panic("apparently this should also be 2")
		}
		val, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			panic("could not parse min value")
		}
		firstIndex := val
		val, err = strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			panic("could not parse max value")
		}
		secondIndex := val

		if p2IsValid(password, []rune(character)[0], firstIndex, secondIndex) {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func Run() {
	part2()
}
