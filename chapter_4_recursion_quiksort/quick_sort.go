package main

import (
	"fmt"
)

func quickSort(slice []int) {
	l := len(slice)
	if l < 2 {
		return
	}
	pivot := l >> 1
	valuePivot := slice[pivot]
	left, right := 0, l-1
	for left <= right {
		for ; slice[left] < valuePivot; left++ {

		}
		for ; slice[right] > valuePivot; right-- {

		}
		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}
	if left >= 0 && left < l {
		quickSort(slice[:left])
		quickSort(slice[left:])
	}
}

func main() {
	slice := []int{2, 5, 3, 23, 56, 1, 65, 0, -1, 8, 2, 1, 7}
	quickSort(slice)
	fmt.Println(slice)
}
