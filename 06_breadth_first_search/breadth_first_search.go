package main

import (
	"fmt"
)

func breadth_first_search(graph map[string][]string, name string) bool {
	queue := []string{name}
	searched := make([]string, 10)

	for len(queue) != 0 {
		person := queue[0]
		queue = queue[1:]

		if includes(person, searched) == false {
			if person_is_seller(person) {
				fmt.Println(person + " is a mango seller!")
				return true
			} else {
				searched = append(searched, person)
				queue = append(queue, graph[person]...)
			}
		}
	}

	return false
}

func includes(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func person_is_seller(person string) bool {
	if person[len(person)-1:] == "m" {
		return true
	}
	return false
}

func main() {
	graph := make(map[string][]string)

	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	breadth_first_search(graph, "you")
}
