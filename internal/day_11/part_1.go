package day11

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	stoneScaner := bufio.NewScanner(strings.NewReader(input))
	expandedStones := make([]string, 0)

	for stoneScaner.Scan() {
		stoneRow := stoneScaner.Text()

		stones := strings.Split(stoneRow, " ")
		for _, stone := range stones {
			expandedStones = append(expandedStones, expandStones(25, stone)...)
		}
	}

	return fmt.Sprint(len(expandedStones))
}

func expandStones(blinks int, stone string) []string {
	if blinks == 0 {
		return append(make([]string, 0), stone)
	} else if stone == "0" {
		return expandStones(blinks-1, "1")
	} else if len(stone)%2 == 0 {

		stoneOne := stone[:len(stone)/2]
		stoneTwo := stone[len(stone)/2:]
		stoneOneValue, _ := strconv.ParseInt(stoneOne, 10, 32)
		stoneTwoValue, _ := strconv.ParseInt(stoneTwo, 10, 32)
		stones := make([]string, 0)

		stones = append(stones, expandStones(blinks-1, fmt.Sprint(stoneOneValue))...)
		stones = append(stones, expandStones(blinks-1, fmt.Sprint(stoneTwoValue))...)

		return stones
	} else {
		stoneValue, _ := strconv.ParseInt(stone, 10, 32)
		return expandStones(blinks-1, fmt.Sprint(stoneValue*2024))
	}

}
