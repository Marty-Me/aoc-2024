package daytwo

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	safeReports := 0
	for scanner.Scan() {
		report := scanner.Text()
		levels := strings.Split(report, " ")

		if !hasSafeLevelsWithDampener(&levels) {
			continue
		}

		safeReports += 1
	}
	return fmt.Sprint(safeReports)
}

func hasSafeLevelsWithDampener(levels *[]string) bool {

	if isDescending(levels) || isAscending(levels) {
		return true
	}

	for i := 0; i < len(*levels); i++ {
		levelsList := (*levels)
		subList := make([]string, 0)
		if i > 0 {
			subList = append(subList, levelsList[:i]...)
		}
		subList = append(subList, levelsList[i+1:]...)

		if isDescending(&subList) || isAscending(&subList) {
			return true
		}
	}
	return false
}

func isDescending(levels *[]string) bool {
	prev, _ := strconv.ParseInt((*levels)[0], 10, 32)

	for _, levelStr := range (*levels)[1:] {
		level, _ := strconv.ParseInt(levelStr, 10, 32)

		if prev > level && (prev-level) <= 3 {
			prev = level
			continue
		}

		return false
	}

	return true
}

func isAscending(levels *[]string) bool {
	prev, _ := strconv.ParseInt((*levels)[0], 10, 32)

	for _, levelStr := range (*levels)[1:] {
		level, _ := strconv.ParseInt(levelStr, 10, 32)

		if prev < level && (prev-level) >= -3 {
			prev = level
			continue
		}
		return false
	}

	return true
}
