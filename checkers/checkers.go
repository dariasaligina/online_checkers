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
	for i := 0; i < 8; i+=1{
		for j:=i%2; j < 8; j+=2{
			if i < 3{
				array[i][j] = 'b'
			} else if i >4{
				array[i][j] = 'w'
			}
		}
	} 
	return board(array)
}

func (b board) String() string{
	var sb strings.Builder
	sb.Grow(170)
	for i:= 0; i < 8; i++{
		sb.WriteByte(byte('0'+8-i))
		sb.WriteByte('|')
		for j:=0; j<8; j++{
			if b[i][j] == 0 { 
				if (i+j)%2 == 1{
					sb.WriteRune(' ')
				} else{
					sb.WriteRune('-')
				}
			} else {
				sb.WriteRune(b[i][j])
			}
			sb.WriteByte('|')
		}
		sb.WriteRune('\n')
	}
	sb.WriteString(" |")
	for letter :='a'; letter<='h'; letter++{
		sb.WriteByte(byte(letter));
		sb.WriteByte('|')
	}
	return sb.String()
}