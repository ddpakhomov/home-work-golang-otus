package main

import "fmt"

func getChessboard(size int) string {
	chessboard := ""
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				chessboard += "# "
			} else {
				chessboard += "  "
			}
		}
		chessboard += "\n"
	}
	return chessboard
}

func main() {
	fmt.Println("Пример шахмотной доски 8х8")
	size := 3

	chessboard := getChessboard(size)
	fmt.Println(chessboard)

	fmt.Println("Можно создать доску любого размера, введите размер доски: ")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println("Ошибка ввода размера: ", err)
		return
	}

	chessboard = getChessboard(size)
	fmt.Println(chessboard)
}
