package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := int64(0)
	mulRegex := regexp.MustCompile(`mul\((\d+),((\d+))\)`)
	for scanner.Scan() {
		instructionLine := scanner.Text()
		matches := mulRegex.FindAllStringSubmatch(instructionLine, -1)

		for _, match := range matches {
			leftMul, _ := strconv.ParseInt(match[1], 10, 32)
			rightMul, _ := strconv.ParseInt(match[2], 10, 32)

			result += leftMul * rightMul
		}
	}
	return fmt.Sprint(result)
}
