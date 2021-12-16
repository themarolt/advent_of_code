package day10

import (
	"aoc2021/libs"
	"fmt"
	"strings"
)

func isOpening(char rune) bool {
	return char == '(' || char == '{' || char == '[' || char == '<'
}

func isClosing(openChar rune, char rune) bool {
	switch openChar {
	case '(':
		return char == ')'
	case '{':
		return char == '}'
	case '[':
		return char == ']'
	case '<':
		return char == '>'
	}

	return false
}

func point(char rune) int {
	switch char {
	case ')':
		return 3
	case '}':
		return 1197
	case ']':
		return 57
	case '>':
		return 25137
	}

	panic("not found")
}

func decode(openCharStack libs.List, loc int, chars []rune) int {
	if loc < len(chars) {
		if openCharStack.Size() == 0 {
			// we are opening
			char := chars[loc]

			if isOpening(char) {
				openCharStack.Push(char)
				return decode(openCharStack, loc+1, chars)
			} else {
				return point(char)
			}
		} else {
			// could be opening
			char := chars[loc]
			if isOpening(char) {
				openCharStack.Push(char)
				return decode(openCharStack, loc+1, chars)
			} else {
				lastOpenChar := openCharStack.Last().Value.(rune)
				if isClosing(lastOpenChar, char) {
					openCharStack.Pop()
					return decode(openCharStack, loc+1, chars)
				} else {
					return point(char)
				}
			}
		}
	} else {
		return 0
	}
}

func part1(lines []string) {
	sum := 0
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		chars := []rune(trimmed)

		list := libs.NewLinkedList()
		lineSum := 0
		if len(chars) > 0 {
			if isOpening(chars[0]) {
				list.Push(chars[0])
				lineSum += decode(list, 1, chars)
			} else {
				lineSum += point(chars[0])
			}
		}

		fmt.Println(line, lineSum, sum)
		sum += lineSum
	}

	fmt.Println(sum)
}

func calcPoints(chars libs.List) int {
	res := 0

	for el := chars.Last(); el != nil; el = el.Prev() {
		res = res * 5
		switch el.Value.(rune) {
		case '(':
			res += 1
		case '[':
			res += 2
		case '{':
			res += 3
		case '<':
			res += 4
		default:
			panic("unknown char")
		}
	}

	return res
}

func part2(lines []string) {
	scores := libs.NewLinkedList()

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		chars := []rune(trimmed)

		list := libs.NewLinkedList()
		if len(chars) > 0 {
			if isOpening(chars[0]) {
				list.Push(chars[0])
				res := decode(list, 1, chars)
				if res == 0 && list.Size() != 0 {
					// incomplete
					scores.Push(calcPoints(list))
				}
			}
		}
	}

	arr := make([]int64, scores.Size())
	i := 0
	for el := scores.First(); el != nil; el = el.Next() {
		arr[i] = int64(el.Value.(int))
		i++
	}

	libs.QuickSort(arr)

	half := len(arr) / 2
	fmt.Println(arr[half])
}

func Run() {
	lines := libs.ReadTxtFileLines("days/day10/input.txt")
	part2(lines)
}
