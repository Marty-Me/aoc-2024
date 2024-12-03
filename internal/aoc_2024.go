package aoc_2024

import (
	"embed"
	"errors"

	day_one "martinvankeulen.nl/aoc-2024/internal/day_1"
	day_two "martinvankeulen.nl/aoc-2024/internal/day_2"
	aoc_arguments "martinvankeulen.nl/aoc-2024/pkg/package_reader"
)

type AdventOfCode2024 struct {
}

func (a *AdventOfCode2024) Solve(input aoc_arguments.AocInputArguments) (*string, error) {
	switch input.Day {
	case 1:
		return dayOne(input.Part)
	case 2:
		return dayTwo(input.Part)

	default:
		return nil, errors.New("selected <day> not implemented yet")
	}
}

//go:embed assets
var puzzleInputsFS embed.FS

func dayOne(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_1_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_one.SolvePartOne(string(p1i))
	} else {
		result = day_one.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayTwo(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_2_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_two.SolvePartOne(string(p1i))
	} else {
		result = day_two.SolvePartTwo(string(p1i))
	}
	return &result, nil
}
