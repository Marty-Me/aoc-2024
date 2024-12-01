package aoc_arguments

import (
	"errors"
	"os"
	"strconv"
)

type AocInputArguments struct {
	Day  uint8
	Part uint8
}

func ParseAocInputArguments() (*AocInputArguments, error) {
	ia := os.Args[1:]

	if len(ia) != 2 {
		return nil, errors.New("invalid input, correct usage: 'aoc24 <day> <part>'")
	}

	dayInput := ia[0]
	partInput := ia[1]

	day, err := strconv.ParseUint(dayInput, 10, 8)

	if err != nil || day == 0 || day > 25 {
		return nil, errors.New("invalid input, <day> should be a round number between 1 and 25")
	}

	part, err := strconv.ParseUint(partInput, 10, 8)

	if err != nil || part == 0 || part > 2 {
		return nil, errors.New("invalid input, <part> should be a round number between 1 and 2")
	}

	return &AocInputArguments{
		Day:  uint8(day),
		Part: uint8(part),
	}, nil
}
