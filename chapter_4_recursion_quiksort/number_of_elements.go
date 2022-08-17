package main

import "fmt"

func numberOfElements(s []int) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	return 1 + numberOfElements(s[:l-1])
}

func main() {
	r := numberOfElements([]int{1, 2, 3, 4})
	fmt.Println(r)
}
