package main

import (
	"fmt"
)

func countdown(number int) {
	if number < 0 {
		return
	}

	fmt.Println(number)
	countdown(number - 1)
}

func main() {
	countdown(5)
}
