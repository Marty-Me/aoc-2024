package day5

import (
	"bufio"
	"fmt"
	"strings"
)

type Vector2D struct {
	x int
	y int
}

type Guard struct {
	position  Vector2D
	direction Vector2D
}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	grid := make([][]rune, 0)
	var guard Guard
	obstaclePositionsX := make(map[int][]int)
	obstaclePositionsY := make(map[int][]int)
	xLen := 0
	yLen := 0

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		xAxis := make([]rune, 0)

		clObstaclePositions := make([]int, 0)
		for x, char := range line {
			xAxis = append(xAxis, char)
			if char == '^' {
				guard = Guard{
					position: Vector2D{
						x: x,
						y: y,
					},
					direction: Vector2D{
						x: 0,
						y: -1,
					},
				}
			}
			if char == '#' {
				clObstaclePositions = append(clObstaclePositions, x)
				_, exists := obstaclePositionsY[x]

				if !exists {
					obstaclePositionsY[x] = make([]int, 0)
				}
				obstaclePositionsY[x] = append(obstaclePositionsY[x], y)
			}
		}
		obstaclePositionsX[y] = clObstaclePositions
		grid = append(grid, xAxis)
		xLen = len(xAxis)
		y += 1
	}
	yLen = len(grid)

	stepCount := trackSteps(guard, obstaclePositionsX, obstaclePositionsY, xLen, yLen)
	return fmt.Sprint(stepCount)
}

func trackSteps(guard Guard, obstaclePositionsX map[int][]int, obstaclePositionsY map[int][]int, xLen int, yLen int) int {
	stepsMade := make(map[Vector2D]interface{})
	stepsMade[guard.position] = nil

	for {
		gdX := guard.direction.x
		gdY := guard.direction.y
		gX := guard.position.x
		gY := guard.position.y

		if gdX > 0 {
			closestX := -1

			for _, obstacleX := range obstaclePositionsX[gY] {
				if obstacleX > gX && (closestX == -1 || obstacleX < closestX) {
					closestX = obstacleX
				}
			}

			if closestX == -1 {
				for x := gX; x < xLen; x++ {
					stepsMade[Vector2D{
						x: x,
						y: gY,
					}] = nil
				}
				break
			}

			for x := gX; x < closestX; x++ {
				stepsMade[Vector2D{
					x: x,
					y: gY,
				}] = nil
			}
			guard.position = Vector2D{
				x: closestX - 1,
				y: gY,
			}
			guard.direction = Vector2D{
				x: 0,
				y: 1,
			}
		} else if gdX < 0 {
			closestX := -1

			for _, obstacleX := range obstaclePositionsX[gY] {
				if obstacleX < gX && (closestX == -1 || obstacleX > closestX) {
					closestX = obstacleX
				}
			}

			if closestX == -1 {
				for x := gX; x >= 0; x-- {
					stepsMade[Vector2D{
						x: x,
						y: gY,
					}] = nil
				}
				break
			}

			for x := gX; x > closestX; x-- {
				stepsMade[Vector2D{
					x: x,
					y: gY,
				}] = nil
			}
			guard.position = Vector2D{
				x: closestX + 1,
				y: gY,
			}
			guard.direction = Vector2D{
				x: 0,
				y: -1,
			}
		} else if gdY < 0 {
			closestY := -1

			for _, obstacleY := range obstaclePositionsY[gX] {
				if obstacleY < gY && (closestY == -1 || obstacleY > closestY) {
					closestY = obstacleY
				}
			}

			if closestY == -1 {
				for y := gY; y >= 0; y-- {
					stepsMade[Vector2D{
						x: gX,
						y: y,
					}] = nil
				}
				break
			}

			for y := gY; y > closestY; y-- {
				stepsMade[Vector2D{
					x: gX,
					y: y,
				}] = nil
			}
			guard.position = Vector2D{
				x: gX,
				y: closestY + 1,
			}
			guard.direction = Vector2D{
				x: 1,
				y: 0,
			}
		} else if gdY > 0 {
			closestY := -1

			for _, obstacleY := range obstaclePositionsY[gX] {
				if obstacleY > gY && (closestY == -1 || obstacleY < closestY) {
					closestY = obstacleY
				}
			}

			if closestY == -1 {
				for y := gY; y < yLen; y++ {
					stepsMade[Vector2D{
						x: gX,
						y: y,
					}] = nil
				}
				break
			}

			for y := gY; y < closestY; y++ {
				stepsMade[Vector2D{
					x: gX,
					y: y,
				}] = nil
			}
			guard.position = Vector2D{
				x: gX,
				y: closestY - 1,
			}
			guard.direction = Vector2D{
				x: -1,
				y: 0,
			}
		}
	}

	return len(stepsMade)
}
