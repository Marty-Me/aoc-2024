package day12

import (
	"bufio"
	"fmt"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	grid, maxX, maxY := parseGrid(scanner)

	knownPatches := make(map[Tile]*interface{})

	totalPrice := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			tile := Tile{x, y}
			_, found := knownPatches[tile]

			if found {
				continue
			}

			totalPrice += findPatches2(tile, grid, &knownPatches, maxX, maxY)
		}
	}

	return fmt.Sprint(totalPrice)
}

func findPatches2(firstTile Tile, grid [][]rune, knownPatches *map[Tile]*interface{}, maxX int, maxY int) int {
	queue := make([]Tile, 0)
	queue = append(queue, firstTile)
	visitedTiles := make(map[Tile]*interface{})
	var charInFocus rune
	perimeter := 0
	area := 0
	directions := getDirections()
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		if n.x > maxX || n.y > maxY || n.x < 0 || n.y < 0 {
			perimeter++
			continue
		}
		tile := Tile{n.x, n.y}

		_, found := visitedTiles[tile]

		if found {
			if grid[tile.y][tile.x] != charInFocus {
				perimeter++
			}
			continue
		}

		visitedTiles[tile] = nil
		charAt := grid[n.y][n.x]

		if charInFocus == 0 {
			charInFocus = grid[n.y][n.x]
		}

		if charAt == charInFocus {
			(*knownPatches)[tile] = nil
			for _, direction := range directions {
				queue = append(queue, Tile{n.x + direction.deltaX, n.y + direction.deltaY})
			}
			area++
		} else {
			perimeter++
		}
	}

	return area * perimeter
}
