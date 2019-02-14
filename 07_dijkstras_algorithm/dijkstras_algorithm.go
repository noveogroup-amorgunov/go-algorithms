package main

import (
	"fmt"
	"math"
)

func contains(value string, list []string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func find_lowest_cost_node(costs map[string]float64, processed []string) string {
	lowest_cost := math.Inf(0)
	lowest_cost_node := ""

	// Go through each node
	for name, cost := range costs {
		// If it's the lowest cost so far and hasn't been processed yet...
		if cost < lowest_cost && contains(name, processed) == false {
			// ... set it as the new lowest-cost node.
			lowest_cost = cost
			lowest_cost_node = name
		}
	}

	return lowest_cost_node
}

func main() {
	// Initialize graph
	graph := make(map[string]map[string]float64)

	graph["start"] = make(map[string]float64)
	graph["start"]["a"] = 6.0
	graph["start"]["b"] = 2.0

	graph["a"] = make(map[string]float64)
	graph["a"]["finish"] = 1.0

	graph["b"] = make(map[string]float64)
	graph["b"]["a"] = 3.0
	graph["b"]["finish"] = 5.0

	graph["finish"] = make(map[string]float64)

	// The costs table
	costs := make(map[string]float64)

	costs["a"] = 6.0
	costs["b"] = 2.0
	costs["finish"] = math.Inf(0)

	// the parents table
	parents := make(map[string]string)

	parents["a"] = "start"
	parents["b"] = "start"
	parents["finish"] = ""

	processed := make([]string, 10)
	node := find_lowest_cost_node(costs, processed)

	for node != "" {
		cost := costs[node]

		// Go through all the neighbors of this node
		neighbors := graph[node]

		for name, value := range neighbors {
			new_cost := cost + value

			// If it's cheaper to get to this neighbor by going through this node
			if costs[name] > new_cost {
				// ... update the cost for this node
				costs[name] = new_cost
				// This node becomes the new parent for this neighbor
				parents[name] = node
			}
		}

		// Mark the node as processed
		processed = append(processed, node)

		// Find the next node to process, and loop
		node = find_lowest_cost_node(costs, processed)
	}

	fmt.Println("Cost from the start to each node:")
	fmt.Println(costs) // { a: 5, b: 2, fin: 6 }
}
