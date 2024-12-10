package day10

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Vector2 struct {
	x int
	y int
}

func SolvePartOne(input string) string {
	topographicMapScanner := bufio.NewScanner(strings.NewReader(input))

	y := 0

	grid := make([][]int64, 0)
	zeroPositions := make([]Vector2, 0)
	xMax := 0

	for topographicMapScanner.Scan() {
		line := topographicMapScanner.Text()
		xAxis := make([]int64, len(line))
		for x, height := range line {
			if height == '0' {
				zeroPositions = append(zeroPositions, Vector2{x, y})
			}
			value, _ := strconv.ParseInt(string(height), 10, 32)
			xAxis[x] = value
		}
		xMax = len(line)
		grid = append(grid, xAxis)
		y++
	}
	yMax := y

	totalTrailheads := 0
	for _, zeroPosition := range zeroPositions {
		foundPeaks := make(map[Vector2]interface{})
		peaksReachable := findPeak(
			grid,
			0,
			zeroPosition.x,
			zeroPosition.y,
			xMax,
			yMax,
			&foundPeaks,
		)

		totalTrailheads += peaksReachable
	}

	return fmt.Sprint(totalTrailheads)
}

func findPeak(grid [][]int64, currentHeight int, x int, y int, xMax int, yMax int, foundPeaks *map[Vector2]interface{}) int {
	directions := []Vector2{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	_, peakAlreadyFound := (*foundPeaks)[Vector2{x, y}]

	if currentHeight == 9 && !peakAlreadyFound {
		(*foundPeaks)[Vector2{x, y}] = nil
		return 1
	}

	peaksFound := 0
	for _, direction := range directions {
		newX := x + direction.x
		newY := y + direction.y

		if newX >= 0 && newX < xMax && newY >= 0 && newY < yMax {
			heightAt := grid[newY][newX]
			if heightAt == int64(currentHeight)+1 {
				peaksFound += findPeak(grid, currentHeight+1, newX, newY, xMax, yMax, foundPeaks)
			}
		}
	}
	return peaksFound
}
