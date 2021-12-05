package day5

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

type point struct {
	x int64
	y int64
}

type line struct {
	a point
	b point
}

func printLine(l line) {
	fmt.Println(l.a.x, ",", l.a.y, " -> ", l.b.x, ",", l.b.y)
}

func part1(lines []line, xSize int64, ySize int64) {
	// prepare grid
	grid := make([][]int64, xSize)
	for i, _ := range grid {
		grid[i] = make([]int64, ySize)
	}

	for _, line := range lines {
		if line.a.x == line.b.x {
			// x line
			// printLine(line)
			// fmt.Println("Is vertical\n")
			lower := line.a.y
			upper := line.b.y

			if lower > upper {
				lower = line.b.y
				upper = line.a.y
			}

			for i := lower; i <= upper; i++ {
				// fmt.Println("Tagging point ", line.a.x, i, "current value", grid[line.a.x][i])
				grid[line.a.x][i] = grid[line.a.x][i] + 1
			}
		} else if line.a.y == line.b.y {
			// y line
			// printLine(line)
			// fmt.Println("Is horizontal\n")
			lower := line.a.x
			upper := line.b.x

			if lower > upper {
				lower = line.b.x
				upper = line.a.x
			}
			for i := lower; i <= upper; i++ {
				grid[i][line.a.y] = grid[i][line.a.y] + 1
				// fmt.Println("Tagging point ", i, line.a.y, "new value", grid[i][line.a.y])
			}
		}

	}

	// count points
	// libs.Print2DArray(grid)

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	fmt.Println("Count", count)
}

func part2(lines []line, xSize int64, ySize int64) {
	// prepare grid
	grid := make([][]int64, xSize)
	for i, _ := range grid {
		grid[i] = make([]int64, ySize)
	}

	for _, line := range lines {
		if line.a.x == line.b.x {
			// x line
			// printLine(line)
			// fmt.Println("Is vertical\n")
			lower := line.a.y
			upper := line.b.y

			if lower > upper {
				lower = line.b.y
				upper = line.a.y
			}

			for i := lower; i <= upper; i++ {
				// fmt.Println("Tagging point ", line.a.x, i, "current value", grid[line.a.x][i])
				grid[line.a.x][i] = grid[line.a.x][i] + 1
			}
		} else if line.a.y == line.b.y {
			// y line
			// printLine(line)
			// fmt.Println("Is horizontal\n")
			lower := line.a.x
			upper := line.b.x

			if lower > upper {
				lower = line.b.x
				upper = line.a.x
			}
			for i := lower; i <= upper; i++ {
				grid[i][line.a.y] = grid[i][line.a.y] + 1
				// fmt.Println("Tagging point ", i, line.a.y, "new value", grid[i][line.a.y])
			}
		} else {
			// check if diagonal
			var part1 int64 = 0
			var part2 int64 = 0

			if line.a.x > line.b.x {
				part1 = line.b.x - line.a.x
			} else {
				part1 = line.a.x - line.b.x
			}

			if line.a.y > line.b.y {
				part2 = line.b.y - line.a.y
			} else {
				part2 = line.a.y - line.b.y
			}

			if part1 == part2 {
				// printLine(line)
				// fmt.Println("Is diagonal\n")
				// diagonal
				x := line.a.x
				y := line.a.y

				xIncrease := line.a.x < line.b.x
				yIncrease := line.a.y < line.b.y
				for {
					grid[x][y] = grid[x][y] + 1

					// fmt.Println("Tagging point ", x, y, "new value", grid[x][y])
					if xIncrease {
						x++
						if x > line.b.x {
							break
						}
					} else {
						x--
						if x < line.b.x {
							break
						}
					}

					if yIncrease {
						y++
						if y > line.b.y {
							break
						}
					} else {
						y--
						if y < line.b.y {
							break
						}
					}
				}
			}
		}

	}

	// count points
	// libs.Print2DArray(grid)

	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}

	fmt.Println("Count", count)
}

func parsePoint(val string) point {
	parts := strings.Split(val, ",")
	x, err1 := strconv.ParseInt(parts[0], 10, 64)
	if err1 != nil {
		panic("could not parse first part of a point")
	}

	y, err2 := strconv.ParseInt(parts[1], 10, 64)
	if err2 != nil {
		panic("could not parse second part of a point")
	}

	return point{x, y}
}

func Run() {
	input := libs.ReadTxtFileLines("days/day5/input.txt")

	lines := make([]line, len(input))

	var maxX int64 = 0
	var maxY int64 = 0

	for i, lineString := range input {
		pointParts := strings.Split(lineString, " -> ")

		a := parsePoint(strings.TrimSpace(pointParts[0]))
		b := parsePoint(strings.TrimSpace(pointParts[1]))

		if a.x > maxX {
			maxX = a.x
		}
		if b.x > maxX {
			maxX = b.x
		}
		if a.y > maxY {
			maxY = a.y
		}
		if b.y > maxY {
			maxY = b.y
		}

		lines[i] = line{a, b}
	}

	fmt.Println("Parsed", len(lines), "lines", "grid size required", maxX+1, "x", maxY+1)

	part2(lines, maxX+1, maxY+1)
}
