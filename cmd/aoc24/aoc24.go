package main

import (
	"fmt"

	aoc_2024 "martinvankeulen.nl/aoc-2024/internal"
	aoc_arguments "martinvankeulen.nl/aoc-2024/pkg/package_reader"
)

func main() {
	aocInput, err := aoc_arguments.ParseAocInputArguments()

	if err != nil {
		panic(err)
	}

	aoc := aoc_2024.AdventOfCode2024{}
	result, _ := aoc.Solve(*aocInput)

	fmt.Println(*result)
}
