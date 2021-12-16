package day16

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
)

func part1(chars []rune) {
	for _, char := range chars {
		num, err := strconv.ParseUint(string(char), 16, 8)
		if err != nil {
			panic(err)
		}
		for bit := 3; bit >= 0; bit-- {
			if libs.IsBitSet(bit, num) {
				fmt.Print(1)
			} else {
				fmt.Print(0)
			}
		}

		fmt.Println()
	}
}

func Run() {
	input := libs.ReadTxtFileLines("days/day16/test_input.txt")

	chars := []rune(input[0])

	part1(chars)
}

/*

0 = 0000
1 = 0001
2 = 0010
3 = 0011
4 = 0100
5 = 0101
6 = 0110
7 = 0111
8 = 1000
9 = 1001
A = 1010
B = 1011
C = 1100
D = 1101
E = 1110
F = 1111

A = 10
B = 11
C = 12
D = 13
E = 14
F = 15

*/
