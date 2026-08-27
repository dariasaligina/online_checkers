package checkers

import (
	"errors"
	"log"
	"strings"
)

type board [8][8]rune

type coordinate [2]int8

type square string

type move []square

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
	c := [2]int8 {int8((*s)[1]-'1'), int8((*s)[0]-'a'),}
	return coordinate(c)
}

func (c *coordinate)Squere() (square, error){
	s := string(rune('a'+c[1])) + string(rune('1'+c[0]))
	sq, err :=  NewSquare(s)
	return sq, err
}

func NewBoard() board{
	array  := [8][8]rune{}
	for i := 0; i < 8; i+=1{
		for j:=i%2; j < 8; j+=2{
			if i < 3{
				array[i][j] = 'w'
			} else if i >4{
				array[i][j] = 'b'
			}
		}
	} 
	return board(array)
}

func (b board) String() string{
	var sb strings.Builder
	sb.Grow(170)
	for i:= 7; i >= 0; i--{
		sb.WriteByte(byte('1'+i))
		sb.WriteByte('|')
		for j:=0; j<8; j++{
			if b[i][j] == 0 && (i+j)%2 == 1{ 
				sb.WriteRune(' ')
			}else if b[i][j] == 0{
				sb.WriteRune('-')
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
func (c *coordinate) isOutsideBoard() bool{
	if c[0] > 7 || c[0]< 0 || c[1]>7 || c[1]<0{
		return true
	}
	return  false
}

func (b *board) isEmpty(c coordinate) bool{
	ans, err := b.GetCoordinate(c)
	if err != nil {
		return false
	}
	return ans == 0
}

func (b *board) GetCoordinate(c coordinate) (rune,error){
	if c.isOutsideBoard() {
		return 0, errors.New("coordinate is outside board")
	}
	return b[c[0]][c[1]], nil
}

func (c coordinate) add(i, j int8) coordinate{
	c[0] += i
	c[1] += j
	return c
}

func (b *board) SimpleMoves(s square) []move{
	coord := s.Coordinate()
	ans := make([]move, 0)
	val, err := b.GetCoordinate(coord)
	if err != nil{
		log.Fatal(err)
	}
	switch val{
	case 'w':{
		if b.isEmpty(coord.add(1,1)){
			coord1 := coord.add(1,1)
			sq1, _ := coord1.Squere()
			ans = append(ans, move{s, sq1})
		}
		if b.isEmpty(coord.add(1,-1)){
			coord1 := coord.add(1,-1)
			sq1, _ := coord1.Squere()
			ans = append(ans, move{s, sq1})
		}
	}
	case 'b':{
		if b.isEmpty(coord.add(-1,1)){
			coord1 := coord.add(-1,1)
			sq1, _ := coord1.Squere()
			ans = append(ans, move{s, sq1})
		}
		if b.isEmpty(coord.add(-1,-1)){
			coord1 := coord.add(-1,-1)
			sq1, _ := coord1.Squere()
			ans = append(ans, move{s, sq1})
		}
		}
	}
	return ans
}
