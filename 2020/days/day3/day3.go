package day3

import (
	"aoc2020/libs"
	"fmt"
)

/*
open squares: .
trees: #
*/

const openSquare rune = '.'
const tree rune = '#'

func countTrees(input []string, xDir int, yDir int) int {
	x := 0
	y := 0

	treeCount := 0
	for y < len(input) {
		y += yDir

		if y >= len(input) {
			break
		}

		x += xDir

		line := input[y]
		char := []rune(line)[x % len(line)]

		if char == tree {
			treeCount += 1
		}
	}

	return treeCount
}

func part1() {
	input := libs.ReadTxtFileLines("days/day3/input.txt")

	treeCount := countTrees(input, 3, 1)

	fmt.Println(treeCount)
}

func part2() {
	input := libs.ReadTxtFileLines("days/day3/input.txt")

	x0 := countTrees(input, 1, 1)
	x1 := countTrees(input, 3, 1)
	x2 := countTrees(input, 5, 1)
	x3 := countTrees(input, 7, 1)
	x4 := countTrees(input, 1, 2)

	fmt.Println(x0 * x1 * x2 * x3 * x4)
}

func Run() {
	part2()
}
