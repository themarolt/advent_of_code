package day4

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

func arrayContains(arr []int64, n int64) bool {
	for _, num := range arr {
		if num == n {
			return true
		}
	}

	return false
}

func isBoardBingo(arr [][]int64, numbers []int64) []int64 {
	if len(numbers) < 5 {
		return nil
	}

	// check rows
	for i := 0; i < 5; i++ {
		row := arr[i]
		bingo := true
		for _, n := range row {
			if !arrayContains(numbers, n) {
				bingo = false
				break
			}
		}

		if bingo {
			return row
		}
	}

	// check columns
	for x := 0; x < 5; x++ {
		column := make([]int64, 5)
		for y := 0; y < 5; y++ {
			column[y] = arr[y][x]
		}

		bingo := true
		for _, n := range column {
			if !arrayContains(numbers, n) {
				bingo = false
				break
			}
		}

		if bingo {
			return column
		}
	}

	return nil
}

func getUnmarkedSum(board [][]int64, numbers []int64) int64 {
	var sum int64 = 0
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			if !arrayContains(numbers, board[x][y]) {
				sum = sum + board[x][y]
			}
		}
	}

	return sum
}

func part1() {
	lines := libs.ReadTxtFileLines("days/day4/input.txt")
	bingoNumbersStrings := strings.Split(lines[0], ",")
	bingoNumbers := make([]int64, len(bingoNumbersStrings))
	for i := 0; i < len(bingoNumbersStrings); i++ {
		val, err := strconv.ParseInt(bingoNumbersStrings[i], 10, 64)
		if err != nil {
			panic("could not convert integer")
		}
		bingoNumbers[i] = val
	}
	fmt.Println("Bingo Numbers:")
	libs.PrintInt64Array(bingoNumbers)

	boardCount := (len(lines) - 1) / 6
	fmt.Println("We have ", boardCount, " boards!")
	boards := make([][][]int64, boardCount)
	for i := range boards {
		boards[i] = make([][]int64, 5)
	}

	// lets parse boards
	boardIndex := 0
	boardLineIndex := 0
	for i := 2; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		if len(trimmed) == 0 {
			boardIndex++
			boardLineIndex = 0
		} else {
			parts := strings.Split(trimmed, " ")
			row := make([]int64, 5)

			index := 0
			for j := 0; j < len(parts); j++ {
				entry := strings.TrimSpace(parts[j])
				if len(entry) == 0 {
					continue
				}

				val, err := strconv.ParseInt(parts[j], 10, 64)
				if err != nil {
					panic(err)
				}
				row[index] = val
				index++
			}
			boards[boardIndex][boardLineIndex] = row
			boardLineIndex++
		}
	}

	// let's print them out
	for i := 0; i < len(boards); i++ {
		data := boards[i]

		libs.PrintInt642DArray(data)
		fmt.Println()
	}

	for x := 4; x < len(bingoNumbers); x++ {
		numbers := bingoNumbers[0 : x+1]

		for _, board := range boards {
			res := isBoardBingo(board, numbers)

			if res != nil {
				sum := getUnmarkedSum(board, numbers)
				fmt.Println(sum, numbers[x], sum*numbers[x])
				return
			}
		}
	}
}

func part2() {
	lines := libs.ReadTxtFileLines("days/day4/input.txt")
	bingoNumbersStrings := strings.Split(lines[0], ",")
	bingoNumbers := make([]int64, len(bingoNumbersStrings))
	for i := 0; i < len(bingoNumbersStrings); i++ {
		val, err := strconv.ParseInt(bingoNumbersStrings[i], 10, 64)
		if err != nil {
			panic("could not convert integer")
		}
		bingoNumbers[i] = val
	}
	fmt.Println("Bingo Numbers:")
	libs.PrintInt64Array(bingoNumbers)

	boardCount := (len(lines) - 1) / 6
	fmt.Println("We have ", boardCount, " boards!")
	boards := make([][][]int64, boardCount)
	for i := range boards {
		boards[i] = make([][]int64, 5)
	}

	// lets parse boards
	boardIndex := 0
	boardLineIndex := 0
	for i := 2; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		if len(trimmed) == 0 {
			boardIndex++
			boardLineIndex = 0
		} else {
			parts := strings.Split(trimmed, " ")
			row := make([]int64, 5)

			index := 0
			for j := 0; j < len(parts); j++ {
				entry := strings.TrimSpace(parts[j])
				if len(entry) == 0 {
					continue
				}

				val, err := strconv.ParseInt(parts[j], 10, 64)
				if err != nil {
					panic(err)
				}
				row[index] = val
				index++
			}
			boards[boardIndex][boardLineIndex] = row
			boardLineIndex++
		}
	}

	boardAlreadyBingo := make([]int64, boardCount)
	for i, _ := range boardAlreadyBingo {
		boardAlreadyBingo[i] = -1
	}
	boardAlreadyBingoCount := 0

	for x := 4; x < len(bingoNumbers); x++ {
		numbers := bingoNumbers[0 : x+1]

		for index, board := range boards {
			if !arrayContains(boardAlreadyBingo, int64(index)) {
				res := isBoardBingo(board, numbers)

				if res != nil {
					fmt.Println("Board #", index+1, " just had a bingo using first ", x+1, " bingo numbers!")
					boardAlreadyBingo[boardAlreadyBingoCount] = int64(index)
					boardAlreadyBingoCount++

					if boardAlreadyBingoCount == len(boards) {
						sum := getUnmarkedSum(board, numbers)
						fmt.Println(sum, numbers[x], sum*numbers[x])
						return
					}
				}
			}
		}
	}
}

func Run() {
	part2()
}
