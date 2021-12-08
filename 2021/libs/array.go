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

// array:
// 5 2 8 4 1 9

// first step: pick pivot (last element in array) = 9
// partition sub array 5 2 8 4 1 according to pivot
// hIndex = 4, lIndex = 0
// 5 < 9 check lIndex++
// 2 < 9 check lIndex++
// 8 < 9 check lIndex++
// 4 < 9 check lIndex++
// 1 < 9 check lIndex++
// lIndex >= hIndex check (5 >= 4)
// no need to swap pivot and lIndex since same index
// we have two sub arrays now:
// [5 2 8 4 1] AND [9]
// second step: do quicksort for first array [5 2 8 4 1]
// pivot = 1, pivotIndex = 4
// partition sub array [5 2 8 4] according to pivot
// hIndex = 3, lIndex = 0
// 5 < 1 false do nothing with lIndex go on with hIndex
// 4 > 1 true hIndex--
// 8 > 1 true hIndex--
// 2 > 1 true hIndex--
// 5 > 1 true hIndex--
// lIndex >= hIndex check (0 >= -1)
// swap lIndex and pivot so we have array [1 2 8 4 5]
