package daytwo

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	safeReports := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		if !hasSafeLevels(&levels) {
			continue
		}

		safeReports += 1
	}
	return fmt.Sprint(safeReports)
}

func hasSafeLevels(levels *[]string) bool {
	prev := int64(-1)
	direction := int64(-1)

	for _, levelStr := range *levels {
		level, _ := strconv.ParseInt(levelStr, 10, 32)

		if prev == -1 {
			first := level
			second, _ := strconv.ParseInt((*levels)[1], 10, 32)

			magnitude := (first - second)

			if magnitude > 0 {
				direction = -1
			} else {
				direction = 1
			}

			prev = level
			continue
		}

		if direction == -1 && -(prev-level) <= -1 && -(prev-level) >= -3 {
			prev = level
			continue
		}

		if direction == 1 && -(prev-level) >= 1 && -(prev-level) <= 3 {
			prev = level
			continue
		}

		return false

	}
	return true
}
