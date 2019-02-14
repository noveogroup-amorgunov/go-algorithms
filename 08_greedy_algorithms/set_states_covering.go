package main

import (
	"fmt"
)

func contains(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func intersection(one []string, two []string) []string {
	result := make([]string, 10)

	for _, value := range one {
		if contains(value, two) {
			result = append(result, value)
		}
	}

	return result
}

func difference(one []string, two []string) []string {
	result := make([]string, 0, 10)

	for _, value := range one {
		if contains(value, two) == false {
			result = append(result, value)
		}
	}

	return result
}

func main() {
	states_needed := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := make(map[string][]string)
	final_stations := make([]string, 0, 10)

	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	for len(states_needed) != 0 {
		best_station := ""
		states_covered := make([]string, 0, 10)

		for station, states := range stations {
			covered := intersection(states_needed, states)

			if len(covered) > len(states_covered) {
				best_station = station
				states_covered = covered
			}
		}

		states_needed = difference(states_needed, states_covered)
		final_stations = append(final_stations, best_station)
	}

	fmt.Println(final_stations) // [kone kfive ktwo kthree]
}
