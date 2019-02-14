package main

import (
	"fmt"
)

func find_smallest(arr []int) int {
	smallest := arr[0]
	smallest_index := 0

	for index, item := range arr {
		if item < smallest {
			smallest = item
			smallest_index = index
		}
	}

	return smallest_index
}

func selection_sort(arr []int) []int {
	sorted_arr := make([]int, len(arr))

	for index := range arr {
		smallest := find_smallest(arr)
		sorted_arr[index] = arr[smallest]

		// Delete smallest value
		// @see https://stackoverflow.com/a/16252034
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return sorted_arr
}

func main() {
	arr := []int{21, 3, 2, 13, 5, 17}

	fmt.Println(selection_sort(arr))
}
