package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	value  string
	blinks int
}

func SolvePartTwo(input string) string {
	stoneScaner := bufio.NewScanner(strings.NewReader(input))
	stoneCount := 0

	cache := make(map[Config]int)

	for stoneScaner.Scan() {
		stoneRow := stoneScaner.Text()
		stones := strings.Split(stoneRow, " ")
		for _, stone := range stones {
			stoneCount += expandStones2(75, stone, &cache)
		}
	}

	return fmt.Sprint(stoneCount)
}

func expandStones2(blinks int, stone string, cache *map[Config]int) int {
	if v, exists := (*cache)[Config{stone, blinks}]; exists {
		return v
	}

	result := 0
	if blinks == 0 {
		result = 1
	} else if stone == "0" {
		result = expandStones2(blinks-1, "1", cache)
	} else if len(stone)%2 == 0 {
		stoneOne := stone[:len(stone)/2]
		stoneTwo := stone[len(stone)/2:]
		stoneOneValue, _ := strconv.ParseInt(stoneOne, 10, 32)
		stoneTwoValue, _ := strconv.ParseInt(stoneTwo, 10, 32)

		one := expandStones2(blinks-1, fmt.Sprint(stoneOneValue), cache)
		two := expandStones2(blinks-1, fmt.Sprint(stoneTwoValue), cache)
		result = one + two
	} else {
		stoneValue, _ := strconv.ParseInt(stone, 10, 32)
		result = expandStones2(blinks-1, fmt.Sprint(stoneValue*2024), cache)
	}

	(*cache)[Config{stone, blinks}] = result
	return result
}
