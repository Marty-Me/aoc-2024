package day8

import (
	"bufio"
	"fmt"
	"strings"
)

type Vector2 struct {
	x int
	y int
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	antennaLocations := make(map[rune][]Vector2)
	antinodeLocations := make(map[Vector2]*interface{})
	maxX := 0

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line)

		for x, value := range line {
			if value != '.' {
				_, found := antennaLocations[value]

				if !found {
					antennaLocations[value] = make([]Vector2, 0)
				}

				antennaLocations[value] = append(antennaLocations[value], Vector2{x: x, y: y})
			}
		}

		y++
	}
	maxY := y

	for _, locations := range antennaLocations {
		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}
				locationOne := locations[i]
				locationTwo := locations[j]

				xD := (locationOne.x - locationTwo.x) * 2
				yD := (locationOne.y - locationTwo.y) * 2

				antinodeX := locationTwo.x + xD
				antinodeY := locationTwo.y + yD

				if antinodeX < maxX && antinodeY < maxY && antinodeX >= 0 && antinodeY >= 0 {
					antinodeLocations[Vector2{antinodeX, antinodeY}] = nil
				}

			}
		}
	}

	return fmt.Sprint(len(antinodeLocations))
}
