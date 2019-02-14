package main

import (
	"fmt"
)

func check_voter(voted map[string]string, name string) {
	if _, exists := voted[name]; exists {
		fmt.Println("kick them out!")
		return
	}

	voted[name] = name
	fmt.Println("let them vote!")
}

func main() {
	voted := make(map[string]string)

	check_voter(voted, "tom")
	check_voter(voted, "mike")
	check_voter(voted, "mike")
}
