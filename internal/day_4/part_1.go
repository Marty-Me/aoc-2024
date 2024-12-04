package day4

import (
	"bufio"
	"fmt"
	"strings"
)

func findWordFromPosition(x int, y int, word string, grid [][]rune) int {
	directions := make([]Direction, 0)

	canGoLeft := x-len(word)+1 >= 0
	canGoRight := x+len(word) <= len(grid[y])
	canGoUp := y-len(word)+1 >= 0
	canGoDown := y+len(word) <= len(grid)

	if canGoDown {
		if canGoRight {
			directions = append(directions, rightDown)
		}
		if canGoLeft {
			directions = append(directions, leftDown)
		}
		directions = append(directions, down)
	}

	if canGoUp {
		if canGoRight {
			directions = append(directions, rightUp)
		}
		if canGoLeft {
			directions = append(directions, leftUp)
		}
		directions = append(directions, up)
	}

	if canGoLeft {
		directions = append(directions, backward)
	}

	if canGoRight {
		directions = append(directions, forward)
	}

	validFound := 0

	for _, direction := range directions {
		for i := 1; i < len(word); i++ {
			xDelta := direction.x * i
			yDelta := direction.y * i

			runeAt := grid[y+yDelta][x+xDelta]

			if runeAt != rune(word[i]) {
				break
			}

			if i == len(word)-1 {
				validFound += 1
			}
		}
	}
	return validFound
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	_ = int64(0)

	grid := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()
		xAxis := make([]rune, 0)

		for _, char := range line {
			xAxis = append(xAxis, char)
		}

		grid = append(grid, xAxis)
	}

	result := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'X' {
				wordsFound := findWordFromPosition(x, y, "XMAS", grid)
				result += wordsFound
			}
		}
	}

	return fmt.Sprint(result)
}
