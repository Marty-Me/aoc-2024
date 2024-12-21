package day_14

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

type Robot struct {
	position Vector2
	velocity Vector2
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	robots := ParseRobots(scanner)
	seconds := 100

	quadrant1 := 0
	quadrant2 := 0
	quadrant3 := 0
	quadrant4 := 0

	for _, r := range robots {
		movedX := (r.velocity.x * seconds)

		areaWidth := 101
		areaHeight := 103

		newX := (r.position.x + movedX) % areaWidth
		movedY := r.velocity.y * seconds
		newY := (r.position.y + movedY) % areaHeight

		if newY < 0 {
			newY = areaHeight + newY
		}

		if newX < 0 {
			newX = areaWidth + newX
		}

		right := newX > ((areaWidth - 1) / 2)
		left := newX < ((areaWidth - 1) / 2)
		up := newY < ((areaHeight - 1) / 2)
		down := newY > ((areaHeight - 1) / 2)

		if left && up {
			quadrant1++
		}
		if right && up {
			quadrant2++
		}

		if right && down {
			quadrant3++
		}
		if left && down {
			quadrant4++
		}
	}
	return fmt.Sprint(quadrant1 * quadrant2 * quadrant3 * quadrant4)
}

func ParseRobots(scanner *bufio.Scanner) []Robot {
	robots := make([]Robot, 0)

	for scanner.Scan() {
		line := scanner.Text()

		pv := strings.Split(line, " ")

		p := pv[0]
		v := pv[1]

		kv := strings.Split(p, "=")
		xy := strings.Split(kv[1], ",")

		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])

		position := Vector2{x, y}

		kv = strings.Split(v, "=")
		xy = strings.Split(kv[1], ",")
		x, _ = strconv.Atoi(xy[0])
		y, _ = strconv.Atoi(xy[1])

		velocity := Vector2{x, y}

		r := Robot{position, velocity}
		robots = append(robots, r)
	}
	return robots
}
