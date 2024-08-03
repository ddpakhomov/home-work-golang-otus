package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := low + (high-low)/2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] < target:
			low = mid + 1
		default:
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13}
	target := 7

	result := binarySearch(arr, target)

	if result != -1 {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, result)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
