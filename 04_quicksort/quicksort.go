package main

import (
	"fmt"
)

func quick_sort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	support := array[0]

	less := make([]int, 0, len(array))
	greater := make([]int, 0, len(array))

	for index := 1; index < len(array); index++ {
		if support > array[index] {
			less = append(less, array[index])
		} else {
			greater = append(greater, array[index])
		}
	}

	return append(
		quick_sort(less),
		append(array[:1], quick_sort(greater)...)...,
	)
}

func main() {
	arr := []int{1, 3, 7, -5}

	fmt.Println(quick_sort(arr))
}
