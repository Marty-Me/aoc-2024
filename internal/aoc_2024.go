package aoc_2024

import (
	"embed"
	"errors"

	day_one "martinvankeulen.nl/aoc-2024/internal/day_1"
	day_ten "martinvankeulen.nl/aoc-2024/internal/day_10"
	day_eleven "martinvankeulen.nl/aoc-2024/internal/day_11"
	day_twelve "martinvankeulen.nl/aoc-2024/internal/day_12"
	day_thirteen "martinvankeulen.nl/aoc-2024/internal/day_13"
	day_fourteen "martinvankeulen.nl/aoc-2024/internal/day_14"
	day_fifteen "martinvankeulen.nl/aoc-2024/internal/day_15"
	day_sixteen "martinvankeulen.nl/aoc-2024/internal/day_16"
	day_seventeen "martinvankeulen.nl/aoc-2024/internal/day_17"
	day_eighteen "martinvankeulen.nl/aoc-2024/internal/day_18"
	day_two "martinvankeulen.nl/aoc-2024/internal/day_2"
	day_three "martinvankeulen.nl/aoc-2024/internal/day_3"
	day_four "martinvankeulen.nl/aoc-2024/internal/day_4"
	day_five "martinvankeulen.nl/aoc-2024/internal/day_5"
	day_six "martinvankeulen.nl/aoc-2024/internal/day_6"
	day_seven "martinvankeulen.nl/aoc-2024/internal/day_7"
	day_eight "martinvankeulen.nl/aoc-2024/internal/day_8"
	day_nine "martinvankeulen.nl/aoc-2024/internal/day_9"

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
	case 3:
		return dayThree(input.Part)
	case 4:
		return dayFour(input.Part)
	case 5:
		return dayFive(input.Part)
	case 6:
		return daySix(input.Part)
	case 7:
		return daySeven(input.Part)
	case 8:
		return dayEight(input.Part)
	case 9:
		return dayNine(input.Part)
	case 10:
		return dayTen(input.Part)
	case 11:
		return dayEleven(input.Part)
	case 12:
		return dayTwelve(input.Part)
	case 13:
		return dayThirteen(input.Part)
	case 14:
		return dayFourteen(input.Part)
	case 15:
		return dayFifteen(input.Part)
	case 16:
		return daySixteen(input.Part)
	case 17:
		return daySeventeen(input.Part)
	case 18:
		return dayEighteen(input.Part)

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

func dayThree(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_3_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_three.SolvePartOne(string(p1i))
	} else {
		result = day_three.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayFour(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_4_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_four.SolvePartOne(string(p1i))
	} else {
		result = day_four.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayFive(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_5_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_five.SolvePartOne(string(p1i))
	} else {
		result = day_five.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func daySix(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_6_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_six.SolvePartOne(string(p1i))
	} else {
		result = day_six.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func daySeven(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_9_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_seven.SolvePartOne(string(p1i))
	} else {
		result = day_seven.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayEight(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_8_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_eight.SolvePartOne(string(p1i))
	} else {
		result = day_eight.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayNine(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_9_p1.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_nine.SolvePartOne(string(p1i))
	} else {
		result = day_nine.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayTen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_10.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_ten.SolvePartOne(string(p1i))
	} else {
		result = day_ten.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayEleven(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_11.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_eleven.SolvePartOne(string(p1i))
	} else {
		result = day_eleven.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayTwelve(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_12.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_twelve.SolvePartOne(string(p1i))
	} else {
		result = day_twelve.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayThirteen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_13.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_thirteen.SolvePartOne(string(p1i))
	} else {
		result = day_thirteen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayFourteen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_14.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_fourteen.SolvePartOne(string(p1i))
	} else {
		result = day_fourteen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayFifteen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_15.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_fifteen.SolvePartOne(string(p1i))
	} else {
		result = day_fifteen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func daySixteen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_16.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_sixteen.SolvePartOne(string(p1i))
	} else {
		result = day_sixteen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func daySeventeen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_17.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_seventeen.SolvePartOne(string(p1i))
	} else {
		result = day_seventeen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}

func dayEighteen(part uint8) (*string, error) {
	p1i, err := puzzleInputsFS.ReadFile("assets/day_18.txt")

	if err != nil {
		panic(err)
	}

	var result string

	if part == 1 {
		result = day_eighteen.SolvePartOne(string(p1i))
	} else {
		result = day_eighteen.SolvePartTwo(string(p1i))
	}
	return &result, nil
}
