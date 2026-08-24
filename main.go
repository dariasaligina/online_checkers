package main

import (
	"fmt"
	"online_checkers/checkers"
)

func main() {
	board := checkers.NewBoard()
	fmt.Println(board)
	fmt.Println(len(board.String()))
}