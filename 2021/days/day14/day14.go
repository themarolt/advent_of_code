package day14

import (
	"aoc2021/libs"
	"fmt"
	"strings"
)

func part1(initial string, mappings map[string]string) {
	// slow
	fmt.Println("Initial:", initial)

	current := initial
	for step := 0; step < 10; step++ {
		fmt.Println("Processing step: ", step+1)
		newPolymer := ""

		for i := 0; i < len(current)-1; i++ {
			substring := current[i : i+2]
			added := false
			to, present := mappings[substring]
			if present {
				newPolymer += substring[0:1] + to
				added = true
			}

			if !added {
				newPolymer += substring[0:1]
			}

			if i == len(current)-2 {
				newPolymer += substring[1:2]
			}
		}

		current = newPolymer
	}

	countMap := make(map[rune]int)

	for _, char := range current {
		val, present := countMap[char]

		if present {
			countMap[char] = val + 1
		} else {
			countMap[char] = 1
		}
	}

	maxChar := '-'
	maxCharCount := 0

	minChar := '-'
	minCharCount := 2147483647
	// found max and min
	for key, element := range countMap {
		if maxCharCount < element {
			maxCharCount = element
			maxChar = key
		}

		if minCharCount > element {
			minCharCount = element
			minChar = key
		}
	}

	fmt.Println("Length: ", len(current))
	fmt.Println("Max Char:", string(maxChar), "Count: ", maxCharCount)
	fmt.Println("Min Char:", string(minChar), "Count: ", minCharCount)
	fmt.Println("Result: ", maxCharCount-minCharCount)
}

func createCount(mappings map[string]string) map[string]int {
	newCount := make(map[string]int)

	for key, _ := range mappings {
		newCount[key] = 0
	}

	return newCount
}

func step(mappings map[string]string, count map[string]int) map[string]int {
	newCount := createCount(mappings)

	for key, value := range count {
		keyChars := []rune(key)
		key1 := string(keyChars[0])
		key2 := string(keyChars[1])
		newCount[key1+mappings[key]] += value
		newCount[mappings[key]+key2] += value
	}

	return newCount
}

func part2(initial string, mappings map[string]string) {
	initialChars := []rune(initial)
	count := createCount(mappings)

	for i := 0; i < len(initial)-1; i++ {
		count[initial[i:i+2]] += 1
	}

	for i := 0; i < 4; i++ {
		count = step(mappings, count)
	}

	// count letters
	letterCount := make(map[rune]int)
	for key, _ := range count {
		for _, char := range key {
			letterCount[char] = 0
		}
	}

	for key, value := range count {
		for _, char := range key {
			letterCount[char] += value
		}
	}

	letterCount[initialChars[0]] += 1
	letterCount[initialChars[len(initialChars)-1]] += 1

	finalCount := make(map[rune]int)
	min := 9223372036854775807
	max := 0
	for key, value := range letterCount {
		newValue := value / 2
		finalCount[key] = newValue

		if max < newValue {
			max = newValue
		}

		if min > newValue {
			min = newValue
		}
	}

	fmt.Println("Diff: ", max-min)
}

func Run() {
	input := libs.ReadTxtFileLines("days/day14/test_input.txt")

	initial := strings.TrimSpace(input[0])

	mappings := make(map[string]string)

	for i := 2; i < len(input); i++ {
		parts := strings.Split(strings.TrimSpace(input[i]), " -> ")
		mappings[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	part2(initial, mappings)
}
