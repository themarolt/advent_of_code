package day5

import (
	"aoc2020/libs"
	"fmt"
)

func part1(input []string) {
	max := 0
	for _, seatString := range input {
		seat := []rune(seatString)

		rowValues := seat[:7]

		// detect row
		row := -1
		rowMin := 0
		rowMax := 127

		for _, val := range rowValues {
			if rowMax-rowMin+1 == 2 {
				if val == 'F' {
					row = rowMin
				} else {
					row = rowMax
				}
				break
			}
			half := ((1 + rowMax - rowMin) / 2)

			if val == 'F' {
				rowMax = rowMax - half
			} else if val == 'B' {
				rowMin = rowMin + half
			} else {
				panic("invalid value for row")
			}
		}

		columnValues := seat[7:]
		column := -1
		columnMin := 0
		columnMax := 7

		for _, val := range columnValues {
			if columnMax-columnMin+1 == 2 {
				if val == 'L' {
					column = columnMin
				} else {
					column = columnMax
				}
				break
			}
			half := ((1 + columnMax - columnMin) / 2)

			if val == 'L' {
				columnMax = columnMax - half
			} else if val == 'R' {
				columnMin = columnMin + half
			} else {
				panic("invalid value for column")
			}
		}

		res := row*8 + column
		if max < res {
			max = res
		}
		fmt.Println(row, column, res)
	}

	fmt.Println(max)
}

func part2(input []string) {
	fmt.Println("max seat id", 127*8+7)
	fmt.Println("tickets", len(input))
	seats := make([]bool, 127*8+7)

	for _, seatString := range input {
		seat := []rune(seatString)

		rowValues := seat[:7]

		// detect row
		row := -1
		rowMin := 0
		rowMax := 127

		for _, val := range rowValues {
			if rowMax-rowMin+1 == 2 {
				if val == 'F' {
					row = rowMin
				} else {
					row = rowMax
				}
				break
			}
			half := ((1 + rowMax - rowMin) / 2)

			if val == 'F' {
				rowMax = rowMax - half
			} else if val == 'B' {
				rowMin = rowMin + half
			} else {
				panic("invalid value for row")
			}
		}

		columnValues := seat[7:]
		column := -1
		columnMin := 0
		columnMax := 7

		for _, val := range columnValues {
			if columnMax-columnMin+1 == 2 {
				if val == 'L' {
					column = columnMin
				} else {
					column = columnMax
				}
				break
			}
			half := ((1 + columnMax - columnMin) / 2)

			if val == 'L' {
				columnMax = columnMax - half
			} else if val == 'R' {
				columnMin = columnMin + half
			} else {
				panic("invalid value for column")
			}
		}

		res := row*8 + column
		// fmt.Println(res)
		seats[res] = true
	}

	for i := 0; i < len(seats); i++ {
		if !seats[i] {
			fmt.Println(i)
		}
	}
}

func Run() {
	input := libs.ReadTxtFileLines("days/day5/input.txt")

	part2(input)
}
