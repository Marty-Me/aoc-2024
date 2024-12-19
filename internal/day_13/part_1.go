package day13

import (
	"bufio"
	"fmt"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))

	prizes :=

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) > 0 {
			lineParts := strings.Split(line, " ")
			identifier := lineParts[0]

			if identifier == "Button" {
				fmt.Println("Button")
			}

			if identifier == "Prize:" {
				fmt.Println("Prize")
			}

		}
	}
	return fmt.Sprint(0)
}
