package main

import "fmt"

func binary_search(slice []int, item int) int {
	low := 0
	high := len(slice) - 1
	mid := (0 + high) / 2
	for low <= high {
		if slice[mid] == item {
			return mid
		} else if slice[mid] < item {
			low = mid + 1
			mid = (low + high) / 2
		} else if slice[mid] > item {
			high = mid - 1
			mid = (low + mid) / 2
		}
	}
	// если элемент не найден возвращаем -1
	return -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := binary_search(slice, 10)
	fmt.Println(r)
}
