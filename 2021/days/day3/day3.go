package day3

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
)

const data = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func binaryStringToInt(binaryString string) int64 {
	number, err := strconv.ParseInt(binaryString, 2, 64)
	if err != nil {
		panic(err)
	}

	return number
}

func part1() {
	lines := libs.ReadTxtFileLines("days/day3/input.txt")
	//lines := strings.Split(data, "\n")
	gammaBits := make([]rune, len([]rune(lines[0])))
	epsilonBits := make([]rune, len([]rune(lines[0])))

	for i := 0; i < len([]rune(lines[0])); i++ {
		countOnes := 0
		countZeros := 0

		for j := 0; j < len(lines); j++ {
			line := []rune(lines[j])
			if line[i] == '0' {
				countZeros++
			} else if line[i] == '1' {
				countOnes++
			} else {
				panic("what else is here?")
			}
		}

		if countOnes > countZeros {
			gammaBits[i] = '1'
			epsilonBits[i] = '0'
		} else {
			gammaBits[i] = '0'
			epsilonBits[i] = '1'
		}
	}

	gammaBitsString := string(gammaBits)
	epsilonBitsString := string(epsilonBits)
	gamma := binaryStringToInt(gammaBitsString)
	epsilon := binaryStringToInt(epsilonBitsString)

	fmt.Println(gammaBitsString, gamma, epsilonBitsString, epsilon, gamma*epsilon)
}

func toRuneArray(list libs.List) [][]rune {
	res := make([][]rune, list.Size())

	if list.Size() == 0 {
		return res
	}

	for i := 0; i < list.Size(); i++ {
		res[i] = list.Get(i).Value.([]rune)
	}

	return res
}

func part2() {
	//lines := strings.Split(data, "\n")
	lines := libs.ReadTxtFileLines("days/day3/input.txt")

	m := map[rune]libs.List{
		'0': libs.NewLinkedList(),
		'1': libs.NewLinkedList(),
	}

	charCount := len([]rune(lines[0]))
	filtered := make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		filtered[i] = []rune(lines[i])
	}

	var oxygenStr []rune

	for bitPos := 0; bitPos < charCount; bitPos++ {
		// go through all filtered lines
		for y := 0; y < len(filtered); y++ {
			chars := filtered[y]
			m[chars[bitPos]].Push(chars)
		}

		zeroList := m['0']
		oneList := m['1']

		if zeroList.Size() > oneList.Size() {
			filtered = toRuneArray(zeroList)
		} else if oneList.Size() > zeroList.Size() {
			filtered = toRuneArray(oneList)
		} else {
			filtered = toRuneArray(oneList)
		}

		if len(filtered) == 1 {
			oxygenStr = filtered[0]
			break
		}

		m = map[rune]libs.List{
			'0': libs.NewLinkedList(),
			'1': libs.NewLinkedList(),
		}
	}

	var co2Str []rune

	filtered = make([][]rune, len(lines))
	for i := 0; i < len(lines); i++ {
		filtered[i] = []rune(lines[i])
	}

	for bitPos := 0; bitPos < charCount; bitPos++ {
		// go through all filtered lines
		for y := 0; y < len(filtered); y++ {
			chars := filtered[y]
			m[chars[bitPos]].Push(chars)
		}

		zeroList := m['0']
		oneList := m['1']

		if zeroList.Size() > oneList.Size() {
			filtered = toRuneArray(oneList)
		} else if oneList.Size() > zeroList.Size() {
			filtered = toRuneArray(zeroList)
		} else {
			filtered = toRuneArray(zeroList)
		}

		if len(filtered) == 1 {
			co2Str = filtered[0]
			break
		}

		m = map[rune]libs.List{
			'0': libs.NewLinkedList(),
			'1': libs.NewLinkedList(),
		}
	}

	parsed, err := strconv.ParseInt(string(oxygenStr), 2, 64)
	if err != nil {
		panic(err)
	}
	oxygen := parsed

	parsed, err = strconv.ParseInt(string(co2Str), 2, 64)
	if err != nil {
		panic(err)
	}
	co2 := parsed
	fmt.Println(oxygen, co2, oxygen*co2)
}

func Run() {
	part2()
}
