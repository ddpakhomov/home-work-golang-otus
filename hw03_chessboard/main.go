package main

import "fmt"

func printChessboard(size int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("Пример шахмотной доски 8х8")
	size := 8

	printChessboard(size)

	fmt.Println("Можно создать доску любого размера, введите размер доски: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Ошибка ввода размера: ", err)
		return
	}

	printChessboard(size)
}
