package main

import "testing"

func TestBingarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13}

	if binarySearch(arr, 9) != 4 {
		t.Errorf("Ошибка! Ожидалось 4")
	}
	if binarySearch(arr, 1) != 0 {
		t.Errorf("Ошибка! Ожидалось 0")
	}
	if binarySearch(arr, 13) != 6 {
		t.Errorf("Ошибка! Ожидалось 6")
	}

	if binarySearch(arr, 8) != -1 {
		t.Errorf("Ошибка! Ожидалось -1")
	}
	if binarySearch(arr, 0) != -1 {
		t.Errorf("Ошибка! Ожидалось -1")
	}
	if binarySearch(arr, 15) != -1 {
		t.Errorf("Ошибка! Ожидалось -1")
	}
}
