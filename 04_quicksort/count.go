package main

import (
	"fmt"
)

func count(array []int) int {
	if len(array) == 0 {
		return 0
	}
	return count(array[1:]) + 1
}

func main() {
	arr := []int{1, 3, 7}

	fmt.Println(count(arr))
}
