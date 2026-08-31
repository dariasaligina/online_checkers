package main

import (
	"fmt"
	"online_checkers/checkers"
)

func main() {
	board := checkers.NewBoard()
	fmt.Println(board)
	fmt.Println(len(board.String()))
	s1, _ := checkers.NewSquare("a3")
	s2, _ := checkers.NewSquare("b6")
	s3, _ := checkers.NewSquare("e3")
	fmt.Println(board.SimpleMoves(s1))
	fmt.Println(board.SimpleMoves(s2))
	fmt.Println(board.SimpleMoves(s3))
}