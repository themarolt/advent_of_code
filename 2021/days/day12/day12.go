package day12

import (
	"aoc2021/libs"
	"fmt"
	"strings"
)

const (
	start = "start"
	end   = "end"
)

type Connection struct {
	from string
	to   string
}

type VisitedSmallCave struct {
	loc   string
	count int
}

func (p *Connection) isFromOrTo(val string) bool {
	return p.from == val || p.to == val
}

func isSmallCave(loc string) bool {
	if loc == end || loc == start {
		return false
	}

	first := []rune(loc)[0]

	return first >= 97 && first <= 122
}

func isLargeCave(loc string) bool {
	first := []rune(loc)[0]

	return first >= 65 && first <= 90
}

func isStart(loc string) bool {
	return loc == start
}

func isEnd(loc string) bool {
	return loc == end
}

func (p *Connection) getDestination(from string) string {
	if p.from == from {
		return p.to
	} else if p.to == from {
		return p.from
	}

	panic("invalid from passed")
}

func haveVisitedAlready1(dest string, visited libs.List) bool {
	for el := visited.First(); el != nil; el = el.Next() {
		if el.Value.(string) == dest {
			return true
		}
	}

	return false
}

func findPath1(loc string, possibleConnections []Connection, visitedSmallCaves libs.List, currentPath libs.List) libs.List {
	list := libs.NewLinkedList()

	if isEnd(loc) {
		list.Push(currentPath)
		return list
	}

	// find possible destinations
	possibleDestinations := libs.NewLinkedList()
	for i := 0; i < len(possibleConnections); i++ {
		pc := possibleConnections[i]
		if pc.isFromOrTo(loc) {
			dest := pc.getDestination(loc)
			if isSmallCave(dest) {
				if !haveVisitedAlready1(dest, visitedSmallCaves) {
					possibleDestinations.Push(dest)
				}
			} else if !isStart(dest) {
				possibleDestinations.Push(dest)
			}
		}
	}
	posDests := make([]string, possibleDestinations.Size())
	i := 0
	for el := possibleDestinations.First(); el != nil; el = el.Next() {
		posDests[i] = el.Value.(string)
		i++
	}

	for el := possibleDestinations.First(); el != nil; el = el.Next() {
		clonedPath := currentPath.Clone()
		dest := el.Value.(string)
		clonedPath.Push(dest)

		clonedVisited := visitedSmallCaves.Clone()
		if isSmallCave(dest) {
			clonedVisited.Push(dest)
		}
		newList := findPath1(dest, possibleConnections, clonedVisited, clonedPath)

		for nel := newList.First(); nel != nil; nel = nel.Next() {
			list.Push(nel.Value)
		}
	}

	return list
}

func findVisited(dest string, visited libs.List) *VisitedSmallCave {
	for el := visited.First(); el != nil; el = el.Next() {
		val := el.Value.(VisitedSmallCave)

		if val.loc == dest {
			return &val
		}
	}

	return nil
}

func findPath2(loc string, possibleConnections []Connection, visitedSmallCaves libs.List, currentPath libs.List) libs.List {
	list := libs.NewLinkedList()

	if isEnd(loc) {
		list.Push(currentPath)
		return list
	}

	// find possible destinations
	possibleDestinations := libs.NewLinkedList()
	for i := 0; i < len(possibleConnections); i++ {
		pc := possibleConnections[i]
		if pc.isFromOrTo(loc) {
			dest := pc.getDestination(loc)
			if isSmallCave(dest) {
				// check if we have already visited one small cave twice
				found := false
				visitCount := 0
				for el := visitedSmallCaves.First(); el != nil; el = el.Next() {
					val := el.Value.(VisitedSmallCave)

					if val.count == 2 {
						found = true
					}

					if val.loc == dest {
						visitCount = val.count
					}
				}

				if !found {
					possibleDestinations.Push(dest)
				} else if visitCount < 1 {
					possibleDestinations.Push(dest)
				}
			} else if !isStart(dest) {
				possibleDestinations.Push(dest)
			}
		}
	}
	posDests := make([]string, possibleDestinations.Size())
	i := 0
	for el := possibleDestinations.First(); el != nil; el = el.Next() {
		posDests[i] = el.Value.(string)
		i++
	}

	for el := possibleDestinations.First(); el != nil; el = el.Next() {
		clonedPath := currentPath.Clone()
		dest := el.Value.(string)
		clonedPath.Push(dest)

		clonedVisited := libs.NewLinkedList()
		found := false
		for vel := visitedSmallCaves.First(); vel != nil; vel = vel.Next() {
			val := vel.Value.(VisitedSmallCave)

			if isSmallCave(dest) && val.loc == dest {
				clonedVisited.Push(VisitedSmallCave{dest, val.count + 1})
				found = true
			} else {
				clonedVisited.Push(VisitedSmallCave{val.loc, val.count})
			}
		}

		if isSmallCave(dest) && !found {
			clonedVisited.Push(VisitedSmallCave{dest, 1})
		}

		newList := findPath2(dest, possibleConnections, clonedVisited, clonedPath)

		for nel := newList.First(); nel != nil; nel = nel.Next() {
			list.Push(nel.Value)
		}
	}

	return list
}

func part1(paths []Connection) {
	initial := libs.NewLinkedList()
	initial.Push("start")
	results := findPath1("start", paths, libs.NewLinkedList(), initial)
	for el := results.First(); el != nil; el = el.Next() {
		path := el.Value.(libs.List)
		i := 0
		for pe := path.First(); pe != nil; pe = pe.Next() {
			fmt.Print(pe.Value)

			if i < path.Size()-1 {
				fmt.Print(",")
			}
			i++
		}

		fmt.Println()
	}
	fmt.Println(results.Size())
}

func part2(paths []Connection) {
	initial := libs.NewLinkedList()
	initial.Push("start")
	results := findPath2("start", paths, libs.NewLinkedList(), initial)
	for el := results.First(); el != nil; el = el.Next() {
		path := el.Value.(libs.List)
		i := 0
		for pe := path.First(); pe != nil; pe = pe.Next() {
			fmt.Print(pe.Value)
			if i < path.Size()-1 {
				fmt.Print(",")
			}
			i++
		}

		fmt.Println()
	}
	fmt.Println(results.Size())
}

func Run() {
	input := libs.ReadTxtFileLines("days/day12/input.txt")

	connections := make([]Connection, len(input))

	for i := 0; i < len(input); i++ {
		parts := strings.Split(strings.TrimSpace(input[i]), "-")
		c := Connection{parts[0], parts[1]}
		connections[i] = c
	}

	part2(connections)
}
