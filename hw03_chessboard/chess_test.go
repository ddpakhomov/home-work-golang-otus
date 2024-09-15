package main

import (
	"fmt"
	"testing"
)

func TestGetChessboard3x3(t *testing.T) {
	size := 3
	expectedOutput := "#   # \n  #   \n#   # \n"

	chessboard := getChessboard(size)

	if chessboard != expectedOutput {
		fmt.Printf("Ожидаемый результат:\n%s\nПолучили:\n%s\n", expectedOutput, chessboard)
	}
}
