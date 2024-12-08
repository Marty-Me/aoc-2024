package day7

import (
	"bufio"
	"fmt"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	antennaLocations := make(map[rune][]Vector2)
	antinodeLocations := make(map[Vector2]*interface{})
	maxX := 0

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line) - 1

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
	maxY := y - 1

	for _, locations := range antennaLocations {

		if len(locations) > 1 {
			for _, location := range locations {
				antinodeLocations[location] = nil
			}
		}

		for i := 0; i < len(locations); i++ {
			for j := 0; j < len(locations); j++ {
				if i == j {
					continue
				}
				locationOne := locations[i]
				locationTwo := locations[j]

				for k := 1; true; k++ {
					xD := (locationOne.x - locationTwo.x)
					yD := (locationOne.y - locationTwo.y)

					antinodeX := locationTwo.x + (xD * k)
					antinodeY := locationTwo.y + (yD * k)

					if antinodeX <= maxX && antinodeY <= maxY && antinodeX >= 0 && antinodeY >= 0 {
						antinodeLocations[Vector2{antinodeX, antinodeY}] = nil
					} else {
						break
					}
				}
			}
		}
	}
	return fmt.Sprint(len(antinodeLocations))
}
