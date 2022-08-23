package breadth_first_search

import "fmt"

func PersonIsSeller(name string) bool {
	lastIndex := len(name) - 1
	return name[lastIndex] == 'm'
}

func SearchSeller(queue []string, graph map[string][]string) bool {
	personSearched := make(map[string]bool)
	for len(queue) > 0 {

		if PersonIsSeller(queue[0]) {
			fmt.Println(queue[0], "is a mango seller!")
			return true
		}
		for _, val := range graph[queue[0]] {
			// смотрим проверяли мы уже данного человека или нет
			if _, ok := personSearched[val]; !ok {
				personSearched[val] = true
				queue = append(queue, val)
			}
		}
		queue = queue[1:]
	}
	return false
}

func main() {
	// создаем граф
	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	graph["bob"] = []string{"anuj", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anuj"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["jonny"] = []string{}

	// создаем очередь
	var queue []string
	queue = append(queue, graph["you"]...)
	findOrNotSeller := SearchSeller(queue, graph)
	fmt.Println(findOrNotSeller)
}
