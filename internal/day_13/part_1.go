package day13

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

type Game struct {
	prize Vector2
	a     Vector2
	b     Vector2
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	games := parseGames(scanner)
	aCost := 3
	bCost := 1

	tokens := 0

	for _, game := range games {
		prizeX := game.prize.x
		prizeY := game.prize.y

		aX := game.a.x
		bX := game.b.x

		maxA := prizeX / aX
		cost := -1
		for a := maxA; a > 0; a-- {
			if ((prizeX - (a * aX)) % bX) == 0 {
				b := (prizeX - (a * aX)) / bX
				if prizeY != (a*game.a.y)+(b*game.b.y) {
					continue
				}
				if cost == -1 || (a*aCost)+(b*bCost) < cost {
					cost = (a * aCost) + (b * bCost)
				}

			}
		}

		if cost > 0 {
			tokens += cost
		}
	}

	return fmt.Sprint(tokens)
}

func parseGames(scanner *bufio.Scanner) []Game {
	games := make([]Game, 0)

	var a Vector2
	var b Vector2

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			lineParts := strings.Split(line, ":")
			identifier := lineParts[0]

			coords := strings.Split(lineParts[1], ",")

			if identifier == "Button A" {
				x, _ := strconv.Atoi(strings.Split(coords[0], "+")[1])
				y, _ := strconv.Atoi(strings.Split(coords[1], "+")[1])
				a = Vector2{x, y}
			}

			if identifier == "Button B" {
				x, _ := strconv.Atoi(strings.Split(coords[0], "+")[1])
				y, _ := strconv.Atoi(strings.Split(coords[1], "+")[1])
				b = Vector2{x, y}
			}

			if identifier == "Prize" {
				x, _ := strconv.Atoi(strings.Split(coords[0], "=")[1])
				y, _ := strconv.Atoi(strings.Split(coords[1], "=")[1])

				games = append(games, Game{
					prize: Vector2{x, y},
					a:     a,
					b:     b,
				})
			}
		}
	}
	return games
}
