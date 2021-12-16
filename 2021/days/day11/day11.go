package day11

import (
	"aoc2021/libs"
	"fmt"
	"strings"
)

func createEmptyArr(xSize, ySize int) [][]int {
	arr := make([][]int, xSize)
	for i := 0; i < 10; i++ {
		arr[i] = make([]int, ySize)
	}

	return arr
}

func octoFlash(i, j, xSize, ySize int, octopuses [][]int, flashed [][]int) {
	if flashed[i][j] == 1 {
		return
	}

	flashed[i][j] = 1

	if i == 0 && j == 0 {
		// upper left corner
		octopuses[i][j+1]++
		octopuses[i+1][j+1]++
		octopuses[i+1][j]++
	} else if i == xSize-1 && j == 0 {
		// upper right corner
		octopuses[i-1][j]++
		octopuses[i-1][j+1]++
		octopuses[i][j+1]++
	} else if i == 0 && j == ySize-1 {
		// lower left corner
		octopuses[i][j-1]++
		octopuses[i+1][j-1]++
		octopuses[i+1][j]++
	} else if i == xSize-1 && j == ySize-1 {
		// lower right corner
		octopuses[i-1][j]++
		octopuses[i-1][j-1]++
		octopuses[i][j-1]++
	} else if i == 0 {
		// left edge
		octopuses[i][j-1]++
		octopuses[i+1][j-1]++
		octopuses[i+1][j]++
		octopuses[i+1][j+1]++
		octopuses[i][j+1]++
	} else if i == xSize-1 {
		// right edge
		octopuses[i][j-1]++
		octopuses[i-1][j-1]++
		octopuses[i-1][j]++
		octopuses[i-1][j+1]++
		octopuses[i][j+1]++
	} else if j == 0 {
		// top edge
		octopuses[i-1][j]++
		octopuses[i-1][j+1]++
		octopuses[i][j+1]++
		octopuses[i+1][j+1]++
		octopuses[i+1][j]++
	} else if j == ySize-1 {
		// bottom edge
		octopuses[i-1][j]++
		octopuses[i-1][j-1]++
		octopuses[i][j-1]++
		octopuses[i+1][j-1]++
		octopuses[i+1][j]++
	} else {
		// middle
		octopuses[i-1][j-1]++
		octopuses[i][j-1]++
		octopuses[i+1][j-1]++
		octopuses[i+1][j]++
		octopuses[i+1][j+1]++
		octopuses[i][j+1]++
		octopuses[i-1][j+1]++
		octopuses[i-1][j]++
	}

	for i := 0; i < xSize; i++ {
		for j := 0; j < ySize; j++ {
			if octopuses[i][j] > 9 {
				octoFlash(i, j, xSize, ySize, octopuses, flashed)
			}
		}
	}
}

func part1(octopuses [][]int, xSize, ySize int) {
	step := 0

	for {
		flashes := 0
		flashed := createEmptyArr(xSize, ySize)

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				octopuses[i][j]++
			}
		}

		for i := 0; i < xSize; i++ {
			for j := 0; j < ySize; j++ {
				if octopuses[i][j] > 9 {
					octoFlash(i, j, xSize, ySize, octopuses, flashed)
				}
			}
		}

		for i := 0; i < xSize; i++ {
			for j := 0; j < ySize; j++ {
				if octopuses[i][j] > 9 {
					octopuses[i][j] = 0
				}
			}
		}

		for i := 0; i < xSize; i++ {
			for j := 0; j < ySize; j++ {
				if flashed[i][j] > 0 {
					flashes++
				}
			}
		}

		if flashes == 100 {
			fmt.Println(step + 1)
			break
		}

		step++
	}
}

func Run() {
	input := libs.ReadTxtFileLines("days/day11/input.txt")

	arr := createEmptyArr(10, 10)

	for i, line := range input {
		chars := []rune(strings.TrimSpace(line))
		for j, char := range chars {
			arr[i][j] = int(char - 48)
		}
	}

	part1(arr, 10, 10)
}
