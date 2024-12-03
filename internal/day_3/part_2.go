package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := int64(0)
	keywordsRegex := regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

	do := true
	for scanner.Scan() {
		instructionLine := scanner.Text()
		matches := keywordsRegex.FindAllStringSubmatch(instructionLine, -1)

		for _, match := range matches {
			if match[0][0:3] == "mul" && do {
				leftMul, _ := strconv.ParseInt(match[1], 10, 32)
				rightMul, _ := strconv.ParseInt(match[2], 10, 32)

				result += leftMul * rightMul
			} else if match[0][0:3] == "don" {
				do = false
			} else if match[0][0:3] == "do(" {
				do = true
			}
		}
	}
	return fmt.Sprint(result)
}
