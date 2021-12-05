package day4

import (
	"aoc2020/libs"
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	birthYear      string = "byr"
	issueYear      string = "iyr"
	expirationYear string = "eyr"
	height         string = "hgt"
	hairColor      string = "hcl"
	eyeColor       string = "ecl"
	passportID     string = "pid"
	countryID      string = "cid"
)

var requiredParts = [...]string{
	birthYear,
	issueYear,
	expirationYear,
	height,
	hairColor,
	eyeColor,
	passportID,
}

var validEyeColor = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

var validation = map[string]func(string) bool{
	birthYear: func(val string) bool {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false
		}

		return parsed >= 1920 && parsed <= 2002
	},
	issueYear: func(val string) bool {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false
		}

		return parsed >= 2010 && parsed <= 2020
	},
	expirationYear: func(val string) bool {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return false
		}

		return parsed >= 2020 && parsed <= 2030
	},
	height: func(val string) bool {
		chars := []rune(val)
		if len(chars) > 2 {
			// take last two chars
			metric := string(chars[len(chars)-2:])
			strValue := string(chars[:len(chars)-2])

			if metric == "cm" {
				parsed, err := strconv.ParseInt(strValue, 10, 64)
				if err != nil {
					return false
				}

				return parsed >= 150 && parsed <= 193
			} else if metric == "in" {
				parsed, err := strconv.ParseInt(strValue, 10, 64)
				if err != nil {
					return false
				}

				return parsed >= 59 && parsed <= 76
			}
		}

		return false
	},
	hairColor: func(val string) bool {
		match, _ := regexp.MatchString("^#[0-9a-f]{6}$", val)
		return match
	},
	eyeColor: func(val string) bool {
		for _, valid := range validEyeColor {
			if val == valid {
				return true
			}
		}
		return false
	},
	passportID: func(val string) bool {
		match, _ := regexp.MatchString("^\\d{9}$", val)
		return match
	},
}

func isPartValidPart2(passLines []string) bool {
	fmt.Println("Checking passport: ")
	libs.PrintStringArray(passLines)

	m := make(map[string]bool)
	for _, reqPart := range requiredParts {
		m[reqPart] = false
	}

	for _, pass := range passLines {
		passportProps := strings.Split(pass, " ")

		for _, property := range passportProps {
			parts := strings.Split(property, ":")
			if len(parts) != 2 {
				panic("this should be 2")
			}

			part := parts[0]

			if part == countryID {
				continue
			}

			present, ok := m[part]
			if !ok {
				panic("what to do here? invalid part I believe...")
			}

			if present {
				panic("part is already present... wat do?")
			}

			validationFunction := validation[part]

			if validationFunction == nil {
				panic("we no have validation")
			}

			if validationFunction(strings.TrimSpace(parts[1])) {
				m[part] = true
			}
		}
	}

	valid := true
	for _, reqPart := range requiredParts {
		if !m[reqPart] {
			valid = false
		}
	}

	if valid {
		fmt.Println("Is valid!")
	} else {
		fmt.Println("Is not valid!")
	}

	fmt.Println()

	return valid
}

func isPassValidPart1(passLines []string) bool {
	fmt.Println("Checking passport: ")
	libs.PrintStringArray(passLines)

	m := make(map[string]bool)
	for _, reqPart := range requiredParts {
		m[reqPart] = false
	}

	for _, pass := range passLines {
		passportProps := strings.Split(pass, " ")

		for _, property := range passportProps {
			parts := strings.Split(property, ":")
			if len(parts) != 2 {
				panic("this should be 2")
			}

			part := parts[0]

			if part == countryID {
				continue
			}

			present, ok := m[part]
			if !ok {
				panic("what to do here? invalid part I believe...")
			}

			if present {
				panic("part is already present... wat do?")
			}

			m[part] = true
		}
	}

	valid := true
	for _, reqPart := range requiredParts {
		if !m[reqPart] {
			valid = false
		}
	}

	if valid {
		fmt.Println("Is valid!")
	} else {
		fmt.Println("Is not valid!")
	}

	fmt.Println()

	return valid
}

func part1() {
	lines := libs.ReadTxtFileLines("days/day4/input.txt")

	count := 0

	passport := list.New()

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			// check valid passport
			passLines := make([]string, passport.Len())
			i := 0
			for el := passport.Front(); el != nil; el = el.Next() {
				passLines[i] = el.Value.(string)
				i++
			}

			if isPassValidPart1(passLines) {
				count++
			}

			// reset list
			passport = list.New()
		} else {
			passport.PushBack(line)
		}
	}

	passLines := make([]string, passport.Len())
	i := 0
	for el := passport.Front(); el != nil; el = el.Next() {
		passLines[i] = el.Value.(string)
		i++
	}

	if isPassValidPart1(passLines) {
		count++
	}

	fmt.Println("Valid passports: ", count)
}

func part2() {
	lines := libs.ReadTxtFileLines("days/day4/input.txt")

	count := 0

	passport := list.New()

	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			// check valid passport
			passLines := make([]string, passport.Len())
			i := 0
			for el := passport.Front(); el != nil; el = el.Next() {
				passLines[i] = el.Value.(string)
				i++
			}

			if isPartValidPart2(passLines) {
				count++
			}

			// reset list
			passport = list.New()
		} else {
			passport.PushBack(line)
		}
	}

	passLines := make([]string, passport.Len())
	i := 0
	for el := passport.Front(); el != nil; el = el.Next() {
		passLines[i] = el.Value.(string)
		i++
	}

	if isPartValidPart2(passLines) {
		count++
	}

	fmt.Println("Valid passports: ", count)
}

func Run() {
	part2()
}
