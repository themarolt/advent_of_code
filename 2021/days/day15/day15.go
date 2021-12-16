package day15

import (
	"aoc2021/libs"
	"container/heap"
	"fmt"
	"math"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    *PointInfo // The value of the item; arbitrary.
	priority int        // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type PointInfo struct {
	prevPoint      *libs.IntPoint
	x, y, distance int
}

func getViableMoves(x, y, lenX, lenY int, input [][]int, visited [][]bool) libs.List {
	possible := libs.NewLinkedList()

	if x == 0 && y == 0 {
		// upper left corner
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
	} else if x == lenX-1 && y == 0 {
		// uper right corner
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
	} else if x == 0 && y == lenY-1 {
		// lower left corner
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
	} else if x == lenX-1 && y == lenY-1 {
		// lower right corner
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
	} else if x == 0 {
		// left edge
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
	} else if x == lenX-1 {
		// right edge
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
	} else if y == 0 {
		// top edge
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
	} else if y == lenY-1 {
		// bottom edge
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
	} else {
		// middle
		if !visited[x][y-1] {
			possible.Push(libs.IntPoint{X: x, Y: y - 1})
		}
		if !visited[x+1][y] {
			possible.Push(libs.IntPoint{X: x + 1, Y: y})
		}
		if !visited[x][y+1] {
			possible.Push(libs.IntPoint{X: x, Y: y + 1})
		}
		if !visited[x-1][y] {
			possible.Push(libs.IntPoint{X: x - 1, Y: y})
		}
	}

	return possible
}

var visitedCount int = 0

func calcDistance(pointInfo, targetPointInfo *PointInfo, move libs.IntPoint, numbers [][]int) bool {
	distance := pointInfo.distance + numbers[move.X][move.Y]
	if targetPointInfo.distance > distance {
		targetPointInfo.prevPoint = new(libs.IntPoint)
		targetPointInfo.prevPoint.X = pointInfo.x
		targetPointInfo.prevPoint.Y = pointInfo.y
		targetPointInfo.distance = distance
		return true
	}

	return false
}

func smallestDistance(lenX, lenY int, distances [][]*PointInfo, visited [][]bool) *PointInfo {
	var pointInfo *PointInfo = nil
	found := false
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			if !visited[x][y] {
				currentPoint := distances[x][y]
				if pointInfo == nil || pointInfo.distance > currentPoint.distance {
					pointInfo = currentPoint
					found = true
				}
			}
		}
	}

	if found {
		return pointInfo
	}

	return nil
}

func distanceBetween(x1, y1, x2, y2 int) int {
	x := float64(x2 - x1)
	y := float64(y2 - y1)

	return int(math.Sqrt((x * x) + (y * y)))
}

func dijsktra(lenX, lenY int, numbers [][]int, distances [][]*PointInfo, visited [][]bool) {
	defer libs.TimeTrack("dijsktra")()

	// find unvisited point with the smallest distance
	pointInfo := smallestDistance(lenX, lenY, distances, visited)

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{value: pointInfo, priority: pointInfo.distance, index: 1})

	// priority queue implementation
	for visitedCount != lenX*lenY {
		if pq.Len() == 0 {
			break
		}

		pointInfo = heap.Pop(&pq).(*Item).value

		if visited[pointInfo.x][pointInfo.y] {
			continue
		}

		moves := getViableMoves(pointInfo.x, pointInfo.y, lenX, lenY, numbers, visited)

		for el := moves.First(); el != nil; el = el.Next() {
			move := el.Value.(libs.IntPoint)

			if visited[move.X][move.Y] {
				continue
			}

			targetPointInfo := distances[move.X][move.Y]

			if calcDistance(pointInfo, targetPointInfo, move, numbers) {
				prio := targetPointInfo.distance
				heap.Push(&pq, &Item{value: targetPointInfo, priority: prio, index: 1})
			}
		}

		visited[pointInfo.x][pointInfo.y] = true
		visitedCount++
	}

	// slow implementation
	// for pointInfo != nil {
	// 	moves := getViableMoves(pointInfo.x, pointInfo.y, lenX, lenY, numbers, visited)

	// 	for el := moves.First(); el != nil; el = el.Next() {
	// 		move := el.Value.(libs.IntPoint)
	// 		targetPointInfo := distances[move.X][move.Y]

	// 		calcDistance(pointInfo, targetPointInfo, move, numbers)
	// 	}

	// 	visited[pointInfo.x][pointInfo.y] = true
	// 	visitedCount++
	// 	if pointInfo.x == lenX-1 && pointInfo.y == lenY-1 {
	// 		break
	// 	}

	// 	pointInfo = smallestDistance(lenX, lenY, distances, visited)
	// }
}

func part1(input [][]int) {
	// dijkstra
	lenX, lenY := len(input), len(input[0])
	distances := make([][]*PointInfo, lenX)
	for x := 0; x < lenX; x++ {
		distances[x] = make([]*PointInfo, lenY)
	}

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			p := new(PointInfo)
			if x == 0 && y == 0 {
				p.distance = 0
			} else {
				p.distance = libs.MaxInt
			}
			p.x = x
			p.y = y
			p.prevPoint = nil
			distances[x][y] = p
		}
	}

	visited := make([][]bool, lenX)
	for x := 0; x < lenX; x++ {
		visited[x] = make([]bool, lenY)
	}

	dijsktra(lenX, lenY, input, distances, visited)

	fmt.Println(distances[lenX-1][lenY-1].distance)
}

const factor = 30

func part2(in [][]int) {
	// build larger array
	origLenX, origLenY := len(in), len(in[0])
	lenX, lenY := origLenX*factor, origLenY*factor
	input := make([][]int, lenX)
	for x := 0; x < lenX; x++ {
		input[x] = make([]int, lenY)
	}

	// first go right factor times
	for fX := 0; fX < factor; fX++ {
		for x := 0; x < origLenX; x++ {
			for y := 0; y < origLenY; y++ {
				targetX := x + (fX * origLenX)

				if targetX < origLenX {
					input[targetX][y] = in[targetX][y]
				} else {
					previousX := x + ((fX - 1) * origLenX)
					prevValue := input[previousX][y]
					newValue := (1 + prevValue) % 9
					if newValue == 0 {
						newValue = 9
					}
					input[targetX][y] = newValue
				}
			}
		}
	}

	// then go down factor times
	for fY := 1; fY < factor; fY++ {
		for x := 0; x < lenX; x++ {
			for y := 0; y < origLenY; y++ {
				targetY := y + (fY * origLenY)
				previousY := y + ((fY - 1) * origLenY)
				prevValue := input[x][previousY]
				newValue := (1 + prevValue) % 9
				if newValue == 0 {
					newValue = 9
				}
				input[x][targetY] = newValue
			}
		}
	}

	// libs.PrintInt2DArray(input)

	distances := make([][]*PointInfo, lenX)
	for x := 0; x < lenX; x++ {
		distances[x] = make([]*PointInfo, lenY)
	}

	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			p := new(PointInfo)
			if x == 0 && y == 0 {
				p.distance = 0
			} else {
				p.distance = libs.MaxInt
			}
			p.x = x
			p.y = y
			p.prevPoint = nil
			distances[x][y] = p
		}
	}

	visited := make([][]bool, lenX)
	for x := 0; x < lenX; x++ {
		visited[x] = make([]bool, lenY)
	}

	dijsktra(lenX, lenY, input, distances, visited)

	fmt.Println(distances[lenX-1][lenY-1].distance)
}

func Run() {
	input := libs.ReadTxtFileLines("days/day15/input.txt")

	xSize := len([]rune(input[0]))
	ySize := len(input)

	arr := make([][]int, xSize)
	for x := 0; x < xSize; x++ {
		arr[x] = make([]int, ySize)
	}

	for y, line := range input {
		chars := []rune(line)
		for x, char := range chars {
			arr[x][y] = libs.ParseRuneNumber(char)
		}
	}

	part2(arr)
}
