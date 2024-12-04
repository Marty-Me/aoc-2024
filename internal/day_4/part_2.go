package day4

import (
	"bufio"
	"fmt"
	"strings"
)

func SolvePartTwo(input string) string {
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
			if grid[y][x] == 'A' {
				wordsFound := findMasFromA(x, y, grid)
				result += wordsFound
			}
		}
	}

	return fmt.Sprint(result)
}

func findMasFromA(x int, y int, grid [][]rune) int {

	opposites := make([]oppositeDirections, 0)

	canGoLeft := x-1 >= 0
	canGoRight := x+1 < len(grid[y])
	canGoUp := y-1 >= 0
	canGoDown := y+1 < len(grid)

	if !canGoDown || !canGoRight || !canGoLeft || !canGoUp {
		return 0
	}

	opposites = append(opposites, oppositeDirections{
		first:  leftUp,
		second: rightDown,
	})

	opposites = append(opposites, oppositeDirections{
		first:  rightUp,
		second: leftDown,
	})

	for _, opposite := range opposites {
		charAtFirst := grid[opposite.first.y+y][opposite.first.x+x]
		charAtSecond := grid[opposite.second.y+y][opposite.second.x+x]

		firstHasSupportedChar := charAtFirst == 'M' || charAtFirst == 'S'
		secondHasSupportedChar := charAtSecond == 'M' || charAtSecond == 'S'
		charsAreDifferent := charAtFirst != charAtSecond
		if !(firstHasSupportedChar && secondHasSupportedChar && charsAreDifferent) {
			return 0
		}
	}
	return 1
}

type oppositeDirections struct {
	first  Direction
	second Direction
}
