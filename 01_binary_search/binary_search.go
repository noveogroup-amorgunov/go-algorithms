package main

import (
	"fmt"
)

func binary_search(sorted_array []int, searchable int) int {
	low_index := 0
	high_index := len(sorted_array) - 1

	for low_index <= high_index {
		middle := (low_index + high_index) / 2
		value := sorted_array[middle]

		if value == searchable {
			return middle
		} else if value < searchable {
			low_index = middle + 1
		} else {
			high_index = middle - 1
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 7, 13, 15, 21}

	fmt.Println(binary_search(arr, 7))
	fmt.Println(binary_search(arr, 2))
}
