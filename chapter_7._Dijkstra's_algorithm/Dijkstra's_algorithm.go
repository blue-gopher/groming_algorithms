package main

import "fmt"

func main() {
	graph := make(map[string]map[string]int)

	graph["start"] = map[string]int{}

	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["fin"] = 1
	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5
	graph["fin"] = map[string]int{}

	// самое большое число
	infinity := 1<<63 - 1

	costs := map[string]int{}
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = infinity

	parents := map[string]string{}
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""

	processed := make(map[string]struct{}, 3)

	// алгоритм Дейкстры
	node := minCost(costs, processed)
	for node != "" {
		cost := costs[node]
		neighbors := graph[node]

		for i, val := range neighbors {
			newCost := cost + val

			if newCost < costs[i] {
				costs[i] = newCost
				parents[i] = node
			}
		}
		processed[node] = struct{}{}
		node = minCost(costs, processed)
	}

	fmt.Println(costs)
}

// нахождение узла с наименьшей стоимостью
func minCost(costs map[string]int, processed map[string]struct{}) string {
	minNode := ""
	minCost := -1

	for node, cost := range costs {
		if _, ok := processed[node]; minCost == -1 && !ok {
			minCost = cost
			minNode = node
			continue
		}

		if _, ok := processed[node]; cost < minCost && !ok {
			minNode = node
			minCost = cost
		}
	}

	return minNode
}
