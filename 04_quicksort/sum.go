package main

import (
	"fmt"
)

func sum(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return array[0] + sum(array[1:])
}

func main() {
	arr := []int{1, 3, 7}

	fmt.Println(sum(arr))
}
