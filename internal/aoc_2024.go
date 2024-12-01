package aoc_2024

import (
	"errors"

	aoc_arguments "martinvankeulen.nl/aoc-2024/pkg/package_reader"
)

type AdventOfCode2024 struct {
}

func (a *AdventOfCode2024) Solve(input aoc_arguments.AocInputArguments) (*string, error) {
	switch input.Day {
	case 1:
		return dayOne(input.Part)

	default:
		return nil, errors.New("selected <day> not implemented yet")
	}
}

func dayOne(part uint8) (*string, error) {
	result := "hello"
	return &result, nil
}
