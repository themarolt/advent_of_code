package day13

import (
	"aoc2021/libs"
	"fmt"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func addIfNotExists(list libs.List, p Point) bool {
	for el := list.First(); el != nil; el = el.Next() {
		pe := el.Value.(Point)

		if pe.x == p.x && pe.y == p.y {
			return false
		}
	}

	list.Push(p)
	return true
}

func part1(initial [][]bool, folds libs.List) {
	points := libs.NewLinkedList()

	lenX, lenY := len(initial), len(initial[0])

	for x := 0; x < len(initial); x++ {
		for y := 0; y < len(initial[0]); y++ {
			if initial[x][y] {
				points.Push(Point{x, y})
			}
		}
	}

	maxX, maxY := 0, 0

	for el := folds.First(); el != nil; el = el.Next() {
		maxX = 0
		maxY = 0
		fold := el.Value.(Point)
		newPoints := libs.NewLinkedList()

		// determine target half
		firstHalf := false
		if fold.x < 0 {
			// fold over y -> x does not change
			if (lenY / 2) <= fold.y {
				firstHalf = true
			}
		} else {
			// fold over x -> y does not change
			if (lenX / 2) <= fold.x {
				firstHalf = true
			}
		}

		for pEl := points.First(); pEl != nil; pEl = pEl.Next() {
			p := pEl.Value.(Point)

			add := false
			newPoint := Point{}

			if fold.x < 0 {
				// fold over y -> x does not change
				if firstHalf {
					if p.y < fold.y {
						// is in the first half
						add = true
						newPoint = Point{x: p.x, y: p.y}
					} else if p.y > fold.y {
						// is in the second half
						// mapping to the first half
						distToFold := p.y - fold.y
						add = true
						newPoint = Point{x: p.x, y: fold.y - distToFold}
					}
				} else {
					if p.y < fold.y {
						// is in the first half
						// mapping to the second half
						distToFold := fold.y - p.y
						// normalize
						add = true
						newPoint = Point{x: p.x, y: distToFold - 1}
					} else if p.y > fold.y {
						// is in the second half
						// normalize
						distToFold := p.y - fold.y
						add = true
						newPoint = Point{x: p.x, y: distToFold - 1}
					}
				}
			} else {
				// fold over x -> y does not change
				if firstHalf {
					if p.x < fold.x {
						// is in the first half
						add = true
						newPoint = Point{x: p.x, y: p.y}
					} else if p.x > fold.x {
						// is in the second half
						// mapping to the first half
						disToFold := p.x - fold.x
						add = true
						newPoint = Point{x: fold.x - disToFold, y: p.y}
					}
				} else {
					// mapping to the second half
					if p.x < fold.x {
						// is in the first half
						// mapping to the second half
						distToFold := fold.x - p.x
						// normalize
						add = true
						newPoint = Point{x: distToFold - 1, y: p.y}
					} else if p.x > fold.x {
						// is in the second half
						// normalize
						distToFold := p.x - fold.x
						add = true
						newPoint = Point{x: distToFold - 1, y: p.y}
					}
				}
			}

			if add {
				added := addIfNotExists(newPoints, newPoint)

				if added {
					if maxX < newPoint.x {
						maxX = newPoint.x
					}
					if maxY < newPoint.y {
						maxY = newPoint.y
					}
				}
			}
		}

		fmt.Println(newPoints.Size())
		points = newPoints
	}

	arr := make([][]bool, maxX+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]bool, maxY+1)
	}

	for el := points.First(); el != nil; el = el.Next() {
		p := el.Value.(Point)

		arr[p.x][p.y] = true
	}

	libs.PrintBool2DArray(arr)
}

func Run() {
	input := libs.ReadTxtFileLines("days/day13/input.txt")

	dotsEnd := false
	points := libs.NewLinkedList()

	foldInstructions := libs.NewLinkedList()
	maxX := 0
	maxY := 0
	for _, line := range input {
		stripped := strings.TrimSpace(line)

		if len(stripped) == 0 {
			dotsEnd = true
			continue
		}

		if !dotsEnd {
			parts := strings.Split(stripped, ",")
			num1, err1 := strconv.ParseInt(parts[0], 10, 32)

			if err1 != nil {
				panic(err1)
			}

			num2, err2 := strconv.ParseInt(parts[1], 10, 32)
			if err2 != nil {
				panic(err2)
			}

			p := Point{x: int(num1), y: int(num2)}

			points.Push(p)

			if p.x > maxX {
				maxX = p.x
			}

			if p.y > maxY {
				maxY = p.y
			}
		}

		if strings.HasPrefix(stripped, "fold along") {
			parts := strings.Split(stripped, "=")
			num, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				panic(err)
			}

			// get dir
			firstPart := []rune(parts[0])
			if firstPart[len(firstPart)-1] == 'x' {
				foldInstructions.Push(Point{x: int(num), y: -1})
			} else if firstPart[len(firstPart)-1] == 'y' {
				foldInstructions.Push(Point{x: -1, y: int(num)})
			} else {
				panic("incorrect format")
			}
		}
	}

	arr := make([][]bool, maxX+1)
	for x := 0; x < maxX+1; x++ {
		arr[x] = make([]bool, maxY+1)
	}

	for el := points.First(); el != nil; el = el.Next() {
		p := el.Value.(Point)

		arr[p.x][p.y] = true
	}

	// libs.PrintBool2DArray(arr)

	fmt.Println("initial size", maxX+1, "x", maxY+1)

	part1(arr, foldInstructions)
}
