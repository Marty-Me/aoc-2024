package dayone

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	teamOneLocationIds := make([]int, 0)
	teamTwoLocationIds := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineLocationIds := strings.Split(line, "   ")

		teamOneLocationId, _ := strconv.ParseInt(lineLocationIds[0], 10, 32)
		teamTwoLocationId, _ := strconv.ParseInt(lineLocationIds[1], 10, 32)

		teamOneLocationIds = append(teamOneLocationIds, int(teamOneLocationId))
		teamTwoLocationIds = append(teamTwoLocationIds, int(teamTwoLocationId))
	}

	sort.Slice(teamOneLocationIds, func(i, j int) bool {
		return teamOneLocationIds[i] < teamOneLocationIds[j]
	})

	sort.Slice(teamTwoLocationIds, func(i, j int) bool {
		return teamTwoLocationIds[i] < teamTwoLocationIds[j]
	})

	result := 0

	for i := range teamOneLocationIds {
		distance := (teamOneLocationIds[i] - teamTwoLocationIds[i])

		if distance < 0 {
			distance = -distance
		}
		result += distance
	}

	return fmt.Sprint(result)
}
