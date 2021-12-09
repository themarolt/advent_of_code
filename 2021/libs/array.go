package libs

import "fmt"

func swapPlaces(arr []int64, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func qsPartition(arr []int64, l, h, pI int) int {
	left := l
	right := h - 1

	pivot := arr[pI]

	for true {
		for arr[left] <= pivot && left <= right {
			left++
		}

		for right > 0 && arr[right] > pivot {
			right--
		}

		if left >= right {
			break
		} else {
			swapPlaces(arr, left, right)
		}
	}

	swapPlaces(arr, left, h)
	return left
}

func quickSort(arr []int64, l, h int) {
	if (h - l) <= 0 {
		return
	}

	pivotIndex := h
	partition := qsPartition(arr, l, h, pivotIndex)
	quickSort(arr, l, partition-1)
	quickSort(arr, partition+1, h)
}

func QuickSort(arr []int64) {
	quickSort(arr, 0, len(arr)-1)
}

func PrintArray(arr []int64) {
	fmt.Print("[")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		if i != len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}

func Print2DArray(arr [][]int64) {
	for y := 0; y < len(arr); y++ {
		PrintArray(arr[y])
	}
}

func UniqueStrings(arr []string) []string {
	list := NewLinkedList()

	for _, val := range arr {
		if !list.Contains(val) {
			list.Push(val)
		}
	}

	res := make([]string, list.Size())
	i := 0
	for el := list.First(); el != nil; el = el.Next() {
		res[i] = el.Value.(string)
		i++
	}

	return res
}

type FilterFunction func(interface{}) bool

func FilterStringArray(arr []string, filter FilterFunction) []string {
	list := NewLinkedList()
	for _, val := range arr {
		if filter(val) {
			list.Push(val)
		}
	}

	res := make([]string, list.Size())
	i := 0
	for el := list.First(); el != nil; el = el.Next() {
		res[i] = el.Value.(string)
		i++
	}

	return res
}

func CloneBoolArray(arr []bool) []bool {
	newArr := make([]bool, len(arr))

	for i, val := range arr {
		newArr[i] = val
	}

	return newArr
}

func BoolArrayCompare(arr1 []bool, arr2 []bool) bool {
	if len(arr1) == len(arr2) {
		for i, val := range arr1 {
			if arr2[i] != val {
				return false
			}
		}

		return true
	} else {
		return false
	}
}

func RuneArrayContains(arr []rune, target rune) bool {
	for _, val := range arr {
		if val == target {
			return true
		}
	}

	return false
}

func RuneArrayAdd(arr []rune, newEl rune) []rune {
	newArr := make([]rune, len(arr)+1)
	for i, e := range arr {
		newArr[i] = e
	}
	newArr[len(arr)] = newEl

	return newArr
}
