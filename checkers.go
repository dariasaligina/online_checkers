package checkers

import "errors"

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

