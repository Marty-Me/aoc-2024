package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := int64(0)
	for scanner.Scan() {
		equation := scanner.Text()
		equationParts := strings.Split(equation, ":")
		expectedAnswer, _ := strconv.ParseInt(equationParts[0], 10, 64)

		equationComponents := strings.Split(equationParts[1][1:], " ")

		perms := generateOptions(len(equationComponents)-1, "*+|")

		for _, options := range perms {
			initialValue, _ := strconv.ParseInt(equationComponents[0], 10, 64)
			partialResult := int64(initialValue)

			for i, permutation := range options {
				a, _ := strconv.ParseInt(equationComponents[i+1], 10, 64)

				if permutation == '*' {
					partialResult *= int64(a)
				}
				if permutation == '+' {
					partialResult += int64(a)
				}

				if permutation == '|' {
					newNumber, _ := strconv.ParseInt(fmt.Sprint(partialResult)+equationComponents[i+1], 10, 64)
					partialResult = newNumber
				}
			}
			if partialResult == expectedAnswer {
				result += int64(expectedAnswer)
				break
			}
		}
	}

	return fmt.Sprint(result)
}
