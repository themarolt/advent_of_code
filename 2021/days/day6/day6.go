package day6

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

func part1(initial []uint8) {
	fishesAtDays := make(map[uint8]int)

	for day := 0; day <= 8; day++ {
		count := 0
		for _, e := range initial {
			if e == uint8(day) {
				count++
			}
		}

		fishesAtDays[uint8(day)] = count
	}

	for day := 0; day < 256; day++ {
		newMap := make(map[uint8]int)

		// move day0 fishes to day6
		// and ad the same amount of fishes do day8
		day0Fishes := fishesAtDays[uint8(0)]
		newMap[uint8(8)] = day0Fishes

		// move the rest down the stream of days
		newMap[uint8(0)] = fishesAtDays[uint8(1)]
		newMap[uint8(1)] = fishesAtDays[uint8(2)]
		newMap[uint8(2)] = fishesAtDays[uint8(3)]
		newMap[uint8(3)] = fishesAtDays[uint8(4)]
		newMap[uint8(4)] = fishesAtDays[uint8(5)]
		newMap[uint8(5)] = fishesAtDays[uint8(6)]
		newMap[uint8(6)] = fishesAtDays[uint8(7)] + day0Fishes
		newMap[uint8(7)] = fishesAtDays[uint8(8)]
		newMap[uint8(8)] = day0Fishes

		fishesAtDays = newMap
	}

	count := 0

	for _, el := range fishesAtDays {
		count = count + el
	}

	fmt.Println(count)
}

func Run() {
	input := libs.ReadTxtFileLines("days/day6/input.txt")
	parts := strings.Split(input[0], ",")
	numbers := make([]uint8, len(parts))
	for i, str := range parts {
		parsed, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic("could not parse a number")
		}

		numbers[i] = uint8(parsed)
	}

	part1(numbers)
}
