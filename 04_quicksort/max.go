package main

import (
	"fmt"
)

func max(array []int) int {
	if len(array) == 1 {
		return array[0]
	}

	max := max(array[1:])
	current := array[0]

	if current > max {
		return current
	}
	return max
}

func main() {
	arr := []int{1, 3, 7, -5}

	fmt.Println(max(arr))
}
