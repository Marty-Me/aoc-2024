package day_15

import (
	"bufio"
	"fmt"
	"strings"
)

type Vector2 struct {
	x int
	y int
}

type Tile = Vector2
type Movement = Vector2

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	grid, actions, startTile, boxPositions := ParseInput(scanner)

	robotPosition := startTile
	for len(*actions) > 0 {
		action := (*actions)[0]
		(*actions) = (*actions)[1:]

		nextTile := Tile{
			robotPosition.x + action.x,
			robotPosition.y + action.y,
		}

		foundSpace := FindSpace(nextTile, action, grid, boxPositions)
		if foundSpace {
			(*grid)[robotPosition] = '.'
			newPosition := Tile{
				robotPosition.x + action.x,
				robotPosition.y + action.y,
			}
			(*grid)[newPosition] = '@'
			robotPosition = newPosition
		}
	}
	boxCoordinatesSum := 0
	for k := range *boxPositions {
		boxCoordinatesSum += (100 * k.y) + k.x
	}

	return fmt.Sprint(boxCoordinatesSum)
}

func FindSpace(tile Tile, movement Movement, grid *map[Tile]rune, boxPositions *map[Tile]*interface{}) bool {
	charAt := (*grid)[tile]

	if charAt == '#' {
		return false
	}

	newTile := Tile{
		tile.x + movement.x,
		tile.y + movement.y,
	}

	if charAt == rune(79) {
		foundSpace := FindSpace(newTile, movement, grid, boxPositions)
		if foundSpace {
			(*grid)[tile] = '.'

			(*grid)[newTile] = charAt
			(*boxPositions)[newTile] = nil
			delete((*boxPositions), tile)
		}
		return foundSpace
	}

	if charAt == '.' {
		return true
	}

	return false
}

func ParseInput(scanner *bufio.Scanner) (*map[Tile]rune, *[]Movement, Tile, *map[Tile]*interface{}) {
	grid := make(map[Tile]rune)
	instructions := make([]Movement, 0)
	boxPositions := make(map[Tile]*interface{})
	var startPosition Tile

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		for x, char := range line {
			if line[0] == '#' {
				tile := Tile{x, y}
				grid[tile] = char
				if char == '@' {
					startPosition = tile
				}
				if char == rune(79) {
					boxPositions[tile] = nil
				}
			} else {
				movement := parseMovement(char)
				instructions = append(instructions, movement)
			}
		}
		y++
	}
	return &grid, &instructions, startPosition, &boxPositions
}

func parseMovement(r rune) Movement {
	if r == '<' {
		return Movement{-1, 0}
	}
	if r == '^' {
		return Movement{0, -1}
	}
	if r == '>' {
		return Movement{1, 0}
	}
	if r == 'v' {
		return Movement{0, 1}
	}
	panic("Undefined character found")
}
