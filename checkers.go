package checkers

import (
	"errors"
	"strings"
)

type board [8][8]rune

type coordinate [2]uint8

type square string

func NewSquare(s string) (square, error) {
	if len(s) != 2 {
		return square(""), errors.New("Invalid coordinate format. Expected 2 characters (e.g., 'e4').")
	}
	if s[0]<'a' || s[0]>'h'{
		return square(""), errors.New("Invalid column. Use letters from 'a' to 'h'.")
	}
	if s[1]<'1' || s[1]>'8'{
		return square(""), errors.New("Invalid row. Use numbers from '1' to '8'")
	}
	return square(s), nil
}

func (s *square)Coordinate() coordinate{
	c := [2]uint8 {(*s)[0]-'a', (*s)[1]-'1'}
	return coordinate(c)
}

func NewBoard() board{
	array  := [8][8]rune{}
	for i := 0; i < 8; i+=2{
		for j:=i%2; j < 8; j+=2{
			if i < 3{
				array[i][j] = 'w'
			} else if i >5{
				array[i][j] = 'b'
			}
		}
	} 
	return board(array)
}

func (b board) String() string{
	var sb strings.Builder
	sb.Grow(64)
	for i:= 0; i < 8; i++{
		for j:=0; j<8; j++{
			sb.WriteRune(b[i][j])
		}
	}
	return sb.String()
}