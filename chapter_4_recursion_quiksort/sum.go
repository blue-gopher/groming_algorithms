package main

import "fmt"

func sum(s []int) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	return s[l-1] + sum(s[:l-1])
}

func main() {
	r := sum([]int{1, 2, 3, 4})
	fmt.Println(r)
}
