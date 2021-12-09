package day8

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

type possibleConfig struct {
	config []bool
	index  int
}

func areCharsUnique(chars []rune) bool {

	for i := 0; i < len(chars)-1; i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				return false
			}
		}
	}

	return true
}

func part1(input []string) {
	count := 0

	for _, line := range input {
		parts := strings.Split(line, "|")

		parts = strings.Split(parts[1], " ")
		fmt.Println(parts)

		for _, output := range parts {
			formatted := strings.TrimSpace(output)
			chars := []rune(formatted)
			length := len(chars)
			if length != 0 && areCharsUnique(chars) {
				if length == 2 {
					fmt.Print(formatted, ", ")
					count++
				} else if length == 3 {
					fmt.Print(formatted, ", ")
					count++
				} else if length == 4 {
					fmt.Print(formatted, ", ")
					count++
				} else if length == 7 {
					fmt.Print(formatted, ", ")
					count++
				}
			}
		}

		fmt.Println()
	}

	fmt.Println(count)
}

func createEmptyPossibilitiesMap() map[rune][]rune {
	possibilities := make(map[rune][]rune)

	possibilities['a'] = make([]rune, 0)
	possibilities['b'] = make([]rune, 0)
	possibilities['c'] = make([]rune, 0)
	possibilities['d'] = make([]rune, 0)
	possibilities['e'] = make([]rune, 0)
	possibilities['f'] = make([]rune, 0)
	possibilities['g'] = make([]rune, 0)

	return possibilities
}

func isDecoded(pos map[rune][]rune) bool {
	for _, list := range pos {
		if len(list) != 1 {
			return false
		}
	}

	return true
}

var zeroSegmentConfig [7]bool = [7]bool{true, true, true, false, true, true, true}
var oneSegmentConfig [7]bool = [7]bool{false, false, true, false, false, true, false}
var twoSegmentConfig [7]bool = [7]bool{true, false, true, true, true, false, true}
var threeSegmentConfig [7]bool = [7]bool{true, false, true, true, false, true, true}
var fourSegmentConfig [7]bool = [7]bool{false, true, true, true, false, true, false}
var fiveSegmentConfig [7]bool = [7]bool{true, true, false, true, false, true, true}
var sixSegmentConfig [7]bool = [7]bool{true, true, false, true, true, true, true}
var sevenSegmentConfig [7]bool = [7]bool{true, false, true, false, false, true, false}
var eightSegmentConfig [7]bool = [7]bool{true, true, true, true, true, true, true}
var nineSegmentConfig [7]bool = [7]bool{true, true, true, true, false, true, true}

func segmentIndex(char rune) int {
	return int(char - 97)
}

func generatePossibleConfigurations(config []bool, processedPossibleLocations libs.List, possibleLocations map[rune]libs.List, allConfigs libs.List) {
	remainingPossibleLocations := make(map[rune]libs.List)
	for key, list := range possibleLocations {
		if !processedPossibleLocations.Contains(key) {
			remainingPossibleLocations[key] = list
		}
	}

	if len(remainingPossibleLocations) == 0 {
		allConfigs.Push(config)
	} else {
		for key, list := range remainingPossibleLocations {
			processedPossibleLocations.Push(key)
			for el := list.First(); el != nil; el = el.Next() {
				clonedConfig := libs.CloneBoolArray(config)
				clonedConfig[el.Value.(int)] = true
				generatePossibleConfigurations(clonedConfig, processedPossibleLocations, possibleLocations, allConfigs)
			}
		}
	}
}

func printConfig(config []bool) {
	if config[0] {
		fmt.Println(" aaaa ")
	} else {
		fmt.Println(" .... ")
	}

	if config[1] && config[2] {
		fmt.Println("b    c")
		fmt.Println("b    c")
	} else if config[1] && !config[2] {
		fmt.Println("b    .")
		fmt.Println("b    .")
	} else if !config[1] && config[2] {
		fmt.Println(".    c")
		fmt.Println(".    c")
	} else {
		fmt.Println(".    .")
		fmt.Println(".    .")
	}

	if config[3] {
		fmt.Println(" dddd ")
	} else {
		fmt.Println(" .... ")
	}

	if config[4] && config[5] {
		fmt.Println("e    f")
		fmt.Println("e    f")
	} else if config[4] && !config[5] {
		fmt.Println("e    .")
		fmt.Println("e    .")
	} else if !config[4] && config[5] {
		fmt.Println(".    f")
		fmt.Println(".    f")
	} else {
		fmt.Println(".    .")
		fmt.Println(".    .")
	}

	if config[6] {
		fmt.Println(" gggg ")
	} else {
		fmt.Println(" .... ")
	}
}

func correctConfigDecode(signal []rune, config []bool, pos map[rune][]rune) {
	for _, char := range signal {
		// find target char
		var targetChar rune
		var arr []rune

		for key, l := range pos {
			if libs.RuneArrayContains(l, char) {
				targetChar = key
				arr = l
				break
			}
		}

		si := segmentIndex(targetChar)
		isPresent := config[si]
		if isPresent && len(arr) > 1 {
			var otherTargetChar rune

			for key, a := range pos {
				if key != targetChar && libs.RuneArrayContains(a, char) {
					otherTargetChar = key
					break
				}
			}

			var otherChar rune
			for _, c := range arr {
				if c != char {
					otherChar = c
					break
				}
			}

			// make sure that otherChar is not in our signal
			if !libs.RuneArrayContains(signal, otherChar) {
				newArr := make([]rune, 1)
				newArr[0] = char
				pos[targetChar] = newArr
				newArr = make([]rune, 1)
				newArr[0] = otherChar
				pos[otherTargetChar] = newArr

				if isDecoded(pos) {
					break
				}
			}
		}
	}
}

func decodeFurther(signal string, pos map[rune][]rune) {
	chars := []rune(signal)

	possibleConfigurations := libs.NewLinkedList()

	config := make([]bool, 7)

	possibleLocations := make(map[rune]libs.List)

	// convert to segments
	for _, char := range chars {
		// find target char
		var targetChar rune
		var list []rune

		for key, l := range pos {
			if libs.RuneArrayContains(l, char) {
				targetChar = key
				list = l
				break
			}
		}

		if len(list) == 1 {
			// perfect already decoded
			config[segmentIndex(targetChar)] = true
		} else if len(list) == 2 {
			firstChar := list[0]
			secondChar := list[1]

			// do we have both segments in this signal
			if libs.RuneArrayContains(chars, firstChar) && libs.RuneArrayContains(chars, secondChar) {
				// find the other target char with the same list
				var otherTargetChar rune

				for key, l := range pos {
					if key != targetChar && libs.RuneArrayContains(l, firstChar) && libs.RuneArrayContains(l, secondChar) {
						otherTargetChar = key
						break
					}
				}

				// mark them both
				config[segmentIndex(targetChar)] = true
				config[segmentIndex(otherTargetChar)] = true
			} else {
				// need to split into multiple possibilities
				// find the other target char with the same char is also inside
				var otherTargetChar rune

				for key, l := range pos {
					if key != targetChar && libs.RuneArrayContains(l, char) {
						otherTargetChar = key
						break
					}
				}

				l := libs.NewLinkedList()
				l.Push(segmentIndex(targetChar))
				l.Push(segmentIndex(otherTargetChar))

				possibleLocations[targetChar] = l
			}
		}
	}

	generatePossibleConfigurations(config, libs.NewLinkedList(), possibleLocations, possibleConfigurations)

	// for key, arr := range pos {
	// 	fmt.Print(string(key) + ": [")
	// 	for i, e := range arr {
	// 		fmt.Print(string(e))

	// 		if i != len(arr)-1 {
	// 			fmt.Print(", ")
	// 		}
	// 	}

	// 	fmt.Print("]\n")
	// }

	// fmt.Println()

	for el := possibleConfigurations.First(); el != nil; el = el.Next() {
		possibleConfig := el.Value.([]bool)
		if libs.BoolArrayCompare(possibleConfig, twoSegmentConfig[:]) {
			// printConfig(possibleConfig)
			correctConfigDecode(chars, possibleConfig, pos)
		}

		if libs.BoolArrayCompare(possibleConfig, threeSegmentConfig[:]) {
			// printConfig(possibleConfig)
			correctConfigDecode(chars, possibleConfig, pos)
		}

		if libs.BoolArrayCompare(possibleConfig, fiveSegmentConfig[:]) {
			// printConfig(possibleConfig)
			correctConfigDecode(chars, possibleConfig, pos)
		}

		if isDecoded(pos) {
			break
		}
	}
}

func decodeNumber(toDecode string, pos map[rune][]rune) rune {
	config := make([]bool, 7)

	for _, char := range []rune(toDecode) {
		// find the correct target char
		for key, arr := range pos {
			if libs.RuneArrayContains(arr, char) {
				config[segmentIndex(key)] = true
				break
			}
		}
	}

	// printConfig(config)

	var toReturn rune = '/'

	if libs.BoolArrayCompare(config, zeroSegmentConfig[:]) {
		toReturn = '0'
	}
	if libs.BoolArrayCompare(config, oneSegmentConfig[:]) {
		toReturn = '1'
	}
	if libs.BoolArrayCompare(config, twoSegmentConfig[:]) {
		toReturn = '2'
	}
	if libs.BoolArrayCompare(config, threeSegmentConfig[:]) {
		toReturn = '3'
	}
	if libs.BoolArrayCompare(config, fourSegmentConfig[:]) {
		toReturn = '4'
	}
	if libs.BoolArrayCompare(config, fiveSegmentConfig[:]) {
		toReturn = '5'
	}
	if libs.BoolArrayCompare(config, sixSegmentConfig[:]) {
		toReturn = '6'
	}
	if libs.BoolArrayCompare(config, sevenSegmentConfig[:]) {
		toReturn = '7'
	}
	if libs.BoolArrayCompare(config, eightSegmentConfig[:]) {
		toReturn = '8'
	}
	if libs.BoolArrayCompare(config, nineSegmentConfig[:]) {
		toReturn = '9'
	}

	// fmt.Println("decoded number: ", string(toReturn))

	if toReturn == '/' {
		panic("what now?")
	}

	return toReturn
}

func part2(input []string) {
	sum := int64(0)
	for _, line := range input {
		pos := createEmptyPossibilitiesMap()

		parts := strings.Split(line, "|")
		signals := libs.FilterStringArray(strings.Split(parts[0], " "), func(el interface{}) bool {
			return len([]rune(el.(string))) != 0
		})

		output := libs.FilterStringArray(strings.Split(parts[1], " "), func(el interface{}) bool {
			return len([]rune(el.(string))) != 0
		})

		// fmt.Println(signals)
		// fmt.Println(output)

		signalForOne := []rune(libs.FilterStringArray(signals, func(i interface{}) bool {
			return len([]rune(i.(string))) == 2
		})[0])

		for _, rune := range signalForOne {
			pos['c'] = libs.RuneArrayAdd(pos['c'], rune)
			pos['f'] = libs.RuneArrayAdd(pos['f'], rune)
		}

		signalForSeven := []rune(libs.FilterStringArray(signals, func(i interface{}) bool {
			return len([]rune(i.(string))) == 3
		})[0])

		for _, rune := range signalForSeven {
			// we can find out signal for d segment
			arr := pos['c']

			if !libs.RuneArrayContains(arr, rune) {
				pos['a'] = libs.RuneArrayAdd(pos['a'], rune)
			}
		}

		signalForFour := []rune(libs.FilterStringArray(signals, func(i interface{}) bool {
			return len([]rune(i.(string))) == 4
		})[0])

		for _, rune := range signalForFour {
			arr := pos['c']

			if !libs.RuneArrayContains(arr, rune) {
				pos['b'] = libs.RuneArrayAdd(pos['b'], rune)
				pos['d'] = libs.RuneArrayAdd(pos['d'], rune)
			}
		}

		signalForEight := []rune(libs.FilterStringArray(signals, func(i interface{}) bool {
			return len([]rune(i.(string))) == 7
		})[0])

		for _, rune := range signalForEight {
			arrC := pos['c']
			arrA := pos['a']
			arrB := pos['b']

			if !libs.RuneArrayContains(arrC, rune) && !libs.RuneArrayContains(arrA, rune) && !libs.RuneArrayContains(arrB, rune) {
				pos['e'] = libs.RuneArrayAdd(pos['e'], rune)
				pos['g'] = libs.RuneArrayAdd(pos['g'], rune)
			}
		}

		signalLength5 := libs.FilterStringArray(signals, func(i interface{}) bool {
			return len([]rune(i.(string))) == 5
		})

		// for key, arr := range pos {
		// 	fmt.Print(string(key) + ": [")
		// 	for i, e := range arr {
		// 		fmt.Print(string(e))

		// 		if i != len(arr)-1 {
		// 			fmt.Print(", ")
		// 		}
		// 	}

		// 	fmt.Print("]\n")
		// }

		// fmt.Println()

		decoded := false
		for !decoded {
			for _, sig := range signalLength5 {
				if isDecoded(pos) {
					decoded = true
					break
				}

				decodeFurther(sig, pos)
			}
		}
		// for key, list := range pos {
		// 	fmt.Print(string(key), ": [")
		// 	for i := 0; i < len(list); i++ {
		// 		if i == len(list)-1 {
		// 			fmt.Print(string(list[i]))
		// 		} else {
		// 			fmt.Print(string(list[i]), ", ")
		// 		}
		// 	}
		// 	fmt.Println("]\n")
		// }

		// lets decode output
		numberStr := make([]rune, len(output))
		for i, toDecode := range output {
			numberStr[i] = decodeNumber(toDecode, pos)
		}

		nr, err := strconv.ParseInt(string(numberStr), 10, 64)

		if err != nil {
			panic(err)
		}
		sum += nr
	}
	fmt.Println(sum)

}

func Run() {
	input := libs.ReadTxtFileLines("days/day8/input.txt")
	part2(input)
}
