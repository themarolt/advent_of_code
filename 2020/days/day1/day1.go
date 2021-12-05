package day1

import (
	"aoc2020/libs"
	"fmt"
	"log"
	"strconv"
)

func part1() {
	input := libs.ReadTxtFileLines("days/day1/input.txt")
	numbers := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		res, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatalf("Error while parsing input line %v with value '%v'. Error: %v", i, input[i], err)
		}

		numbers[i] = res
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			sum := numbers[i] + numbers[j]
			if sum == 2020 {
				fmt.Printf("numbers[%v](%v) + numbers[%v](%v) = %v\n", i, numbers[i], j, numbers[j], sum)
				fmt.Println(numbers[i] * numbers[j])
				return
			}
		}
	}
}

func part2() {
	input := libs.ReadTxtFileLines("days/day1/input.txt")
	numbers := make([]int, len(input))

	for i := 0; i < len(input); i++ {
		res, err := strconv.Atoi(input[i])
		if err != nil {
			log.Fatalf("Error while parsing input line %v with value '%v'. Error: %v", i, input[i], err)
		}

		numbers[i] = res
	}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k<len(numbers); k++ {
				sum := numbers[i] + numbers[j] + numbers[k]
				if sum == 2020 {
					fmt.Printf("numbers[%v](%v) + numbers[%v](%v) + numbers[%v](%v) = %v\n", i, numbers[i], j, numbers[j], k, numbers[k], sum)
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
					return
				}
			}
		}
	}
}

func Run() {
	part2()
}
