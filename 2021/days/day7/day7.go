package day7

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

func mean(nums []int32) int32 {
	var sum int32 = 0
	for _, num := range nums {
		sum += num
	}

	return sum / int32(len(nums))
}

func median(nums []int32) int32 {
	total := len(nums)
	ints := make([]int64, total)
	for i, n := range nums {
		ints[i] = int64(n)
	}

	libs.QuickSort(ints)

	if total%2 == 0 {
		half := total / 2
		return int32(ints[half])
	} else {
		half := (total + 1) / 2
		return int32(ints[half])
	}
}

func part1(positions []int32) {
	// find median
	median := median(positions)
	fmt.Println("Median", median)

	// calc fuel consumption
	total := int32(0)
	for _, val := range positions {
		diff := median - val
		if diff < 0 {
			total += -diff
		} else {
			total += diff
		}
	}

	fmt.Println("Total fuel", total)
}

func triangled(n int32) int32 {
	sum := int32(0)

	for i := int32(1); i <= n; i++ {
		sum += i
	}

	return sum
}

func part2CalcFuel(positions []int32, target int32) int32 {
	total := int32(0)

	for _, pos := range positions {
		if pos > target {
			total += triangled(pos - target)
		} else if pos < target {
			total += triangled(target - pos)
		}
	}

	return total
}

func max(nums []int32) int32 {
	max := nums[0]

	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}

func part2(positions []int32) {
	// find mean
	//mean := mean(positions)
	// fmt.Println("Mean", mean)

	max := max(positions)

	minFuel := int32(-1)

	for i := int32(0); i <= max; i++ {
		fuel := part2CalcFuel(positions, i)
		if minFuel >= 0 {
			if minFuel > fuel {
				minFuel = fuel
			}
		} else {
			minFuel = fuel
		}
	}
	fmt.Println("Total fuel", minFuel)
}

func Run() {
	input := libs.ReadTxtFileLines("days/day7/input.txt")
	split := strings.Split(strings.TrimSpace(input[0]), ",")

	positions := make([]int32, len(split))
	for i, str := range split {
		parsed, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}

		positions[i] = int32(parsed)
	}

	part2(positions)
}
