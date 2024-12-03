package dayone

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	teamOneLocationIds := make([]int, 0)

	locationIdOccuranceCounter := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		lineLocationIds := strings.Split(line, "   ")

		teamOneLocationId, _ := strconv.ParseInt(lineLocationIds[0], 10, 32)
		teamTwoLocationId, _ := strconv.ParseInt(lineLocationIds[1], 10, 32)

		_, keyExists := locationIdOccuranceCounter[int(teamTwoLocationId)]

		if keyExists {
			locationIdOccuranceCounter[int(teamTwoLocationId)] += 1
		} else {
			locationIdOccuranceCounter[int(teamTwoLocationId)] = 1
		}

		teamOneLocationIds = append(teamOneLocationIds, int(teamOneLocationId))
	}

	result := 0
	for _, locationId := range teamOneLocationIds {
		occurance := locationIdOccuranceCounter[locationId]
		result += locationId * occurance
	}

	return fmt.Sprint(result)
}
