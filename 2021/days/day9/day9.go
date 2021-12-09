package day9

import (
	"aoc2021/libs"
	"fmt"
	"strings"
)

func part1(arr [][]int) {
	list := libs.NewLinkedList()

	xSize := len(arr)
	ySize := len(arr[0])

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			val := arr[x][y]
			if y == 0 && x == 0 {
				// upper left corner
				if val < arr[x+1][y] && val < arr[x][y+1] {
					list.Push(val)
				}
			} else if y == 0 && x == (xSize-1) {
				// upper right corner
				if val < arr[x-1][y] && val < arr[x][y+1] {
					list.Push(val)
				}
			} else if y == (ySize-1) && x == 0 {
				// bottom left corner
				if val < arr[x][y-1] && val < arr[x+1][y] {
					list.Push(val)
				}
			} else if y == (ySize-1) && x == (xSize-1) {
				// bottom right corner
				if val < arr[x-1][y] && val < arr[x][y-1] {
					list.Push(val)
				}
			} else if y == 0 {
				// top edge
				if val < arr[x-1][y] && val < arr[x][y+1] && val < arr[x+1][y] {
					list.Push(val)
				}
			} else if y == (ySize)-1 {
				// bottom edge
				if val < arr[x-1][y] && val < arr[x][y-1] && val < arr[x+1][y] {
					list.Push(val)
				}
			} else if x == 0 {
				// left edge
				if val < arr[x][y-1] && val < arr[x+1][y] && val < arr[x][y+1] {
					list.Push(val)
				}
			} else if x == (xSize)-1 {
				// right edge
				if val < arr[x][y-1] && val < arr[x-1][y] && val < arr[x][y+1] {
					list.Push(val)
				}
			} else {
				// center
				if val < arr[x][y-1] && val < arr[x-1][y] && val < arr[x][y+1] && val < arr[x+1][y] {
					list.Push(val)
				}
			}
		}
	}

	sum := 0
	for el := list.First(); el != nil; el = el.Next() {
		sum += el.Value.(int) + 1
	}

	fmt.Println(sum)
}

type point struct {
	x int
	y int
}

func calcBasinSize(p point, xSize int, ySize int, arr [][]int, alreadyProcessed libs.List) int {
	x := p.x
	y := p.y

	if arr[x][y] == 9 {
		return 0
	}

	for el := alreadyProcessed.First(); el != nil; el = el.Next() {
		pp := el.Value.(point)
		if pp.x == x && pp.y == y {
			return 0
		}
	}

	// check surrounding basin
	sum := 1
	alreadyProcessed.Push(p)

	if y == 0 && x == 0 {
		// upper left corner
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
	} else if y == 0 && x == (xSize-1) {
		// upper right corner
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
	} else if y == (ySize-1) && x == 0 {
		// bottom left corner
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
	} else if y == (ySize-1) && x == (xSize-1) {
		// bottom right corner
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
	} else if y == 0 {
		// top edge
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
	} else if y == (ySize)-1 {
		// bottom edge
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
	} else if x == 0 {
		// left edge
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
	} else if x == (xSize)-1 {
		// right edge
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
	} else {
		// center
		sum += calcBasinSize(point{x, y - 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x + 1, y}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x, y + 1}, xSize, ySize, arr, alreadyProcessed)
		sum += calcBasinSize(point{x - 1, y}, xSize, ySize, arr, alreadyProcessed)
	}

	return sum
}

func part2(arr [][]int) {
	processedPoints := libs.NewLinkedList()
	xSize := len(arr)
	ySize := len(arr[0])

	basinSizes := libs.NewLinkedList()

	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			res := calcBasinSize(point{x, y}, xSize, ySize, arr, processedPoints)
			if res > 0 {
				basinSizes.Push(res)
			}
		}
	}

	// determine three largest basins
	basins := make([]int64, basinSizes.Size())
	for i := 0; i < basinSizes.Size(); i++ {
		basins[i] = int64(basinSizes.Get(i).Value.(int))
	}

	libs.QuickSort(basins)

	fmt.Println(basins[len(basins)-1] * basins[len(basins)-2] * basins[len(basins)-3])
}

func Run() {
	input := libs.ReadTxtFileLines("days/day9/input.txt")
	ySize := len(input)
	xSize := len([]rune(input[0]))

	array := make([][]int, xSize)
	for i := 0; i < xSize; i++ {
		array[i] = make([]int, ySize)
	}

	for y, line := range input {
		chars := []rune(strings.TrimSpace(line))
		for x, char := range chars {
			if char >= 48 && char <= 57 {
				array[x][y] = int(char - 48)
			} else {
				panic("unknown char")
			}
		}
	}

	part2(array)
}
