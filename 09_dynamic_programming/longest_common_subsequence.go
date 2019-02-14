package main

import (
	"fmt"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Levenshtein distance
func longest_common_subsequence(word_a, word_b string) (int, [][]int) {
	cells := make([][]int, len(word_a))
	max_value := 0

	for index := range cells {
		cells[index] = make([]int, len(word_b))
	}

	for i := 0; i < len(word_a); i++ {
		for j := 0; j < len(word_b); j++ {
			// The letters match
			if word_a[i] == word_b[j] {
				if i > 0 && j > 0 {
					cells[i][j] = cells[i-1][j-1] + 1
				} else {
					cells[i][j] = 1
				}
			} else {
				// The letters don't match.
				if i > 0 && j > 0 {
					cells[i][j] = max(cells[i-1][j], cells[i][j-1])
				} else if i > 0 {
					cells[i][j] = cells[i-1][j]
				} else if j > 0 {
					cells[i][j] = cells[i][j-1]
				}
			}

			if cells[i][j] > max_value {
				max_value = cells[i][j]
			}
		}
	}

	return max_value, cells
}

func print_result(cells [][]int, word_a, word_b string) {
	splitted_word_a := strings.Split(word_a, "")

	fmt.Println(" ", strings.Split(word_b, ""))
	for index, item := range cells {
		fmt.Println(splitted_word_a[index], item)
	}
}

func main() {
	word_a := "fosh"
	word_b := "fish"

	longest, cells := longest_common_subsequence(word_a, word_b)

	fmt.Println("Longest common subsequence:", longest) // 3
	print_result(cells, word_a, word_b)

	// [1 1 1 1]
	// [1 1 1 1]
	// [1 1 2 2]
	// [1 1 2 3]
}
