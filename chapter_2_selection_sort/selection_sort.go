package main

//поиска наименьшего элемента массива
func FindMin(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if smallest > arr[i] {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

// функция сортировки выбором
func SelectionSort(arr []int) []int {
	l := len(arr)
	newArr := make([]int, l)
	for i := 0; i < l; i++ {
		smallestIndex := FindMin(arr)
		newArr[i] = arr[smallestIndex]
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}
	return newArr
}
