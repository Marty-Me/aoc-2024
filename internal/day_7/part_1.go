package day7

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := int64(0)
	for scanner.Scan() {
		equation := scanner.Text()
		equationParts := strings.Split(equation, ":")
		expectedAnswer, _ := strconv.ParseInt(equationParts[0], 10, 64)

		equationComponents := strings.Split(equationParts[1][1:], " ")

		perms := generateOptions(len(equationComponents)-1, "*+")

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
			}
			if partialResult == expectedAnswer {
				result += int64(expectedAnswer)
				break
			}
		}
	}

	return fmt.Sprint(result)
}

func generateOptions(size int, options string) [][]rune {
	runeOptions := make([][]rune, int(math.Pow(float64(len(options)), float64(size))))

	var generateRecursively func(current string)
	nOptionsGenerated := 0
	generateRecursively = func(current string) {
		if len(current) == size {
			runeOptions[nOptionsGenerated] = []rune(current)
			nOptionsGenerated++
			return
		}

		for _, r := range options {
			generateRecursively(current + string(r))
		}
	}

	generateRecursively("")
	return runeOptions
}
