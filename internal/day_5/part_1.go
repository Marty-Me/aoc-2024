package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Set map[int64]struct{}

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	pageOrderingRules, updates := parsePuzzleInput(scanner)
	result := int64(0)

	for _, update := range updates {
		if validateUpdate(pageOrderingRules, update) {
			result += update[len(update)/2]
		}
	}

	return fmt.Sprint(result)
}

func isPageStillAllowed(page int64, noLongerAllowedPages Set) (bool, int64) {
	for noLongerAllowedPage := range noLongerAllowedPages {
		if noLongerAllowedPage == page {
			return false, page
		}
	}
	return true, -1
}

func parsePageOrderingRules(line string, orderingRules *map[int64]Set) {
	rule := strings.Split(line, "|")
	before, _ := strconv.ParseInt(rule[0], 10, 32)
	after, _ := strconv.ParseInt(rule[1], 10, 32)

	pageNumbers, keyExists := (*orderingRules)[after]

	if !keyExists {
		pageNumbers = make(Set)
	}

	pageNumbers[before] = struct{}{}
	(*orderingRules)[after] = pageNumbers
}

func validateUpdate(pageOrderingRules map[int64]Set, update []int64) bool {
	noLongerAllowedPages := make(Set)

	for _, page := range update {
		stillAllowed, _ := isPageStillAllowed(page, noLongerAllowedPages)
		if stillAllowed {
			shouldBeBefore := pageOrderingRules[page]
			for beforePage := range shouldBeBefore {
				noLongerAllowedPages[beforePage] = struct{}{}
			}
		} else {
			return false
		}
	}
	return true
}

func parsePuzzleInput(scanner *bufio.Scanner) (map[int64]Set, [][]int64) {
	pageOrderingRulesParsed := false
	pageOrderingRules := make(map[int64]Set)
	updates := make([][]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if !pageOrderingRulesParsed {
			if line == "" {
				pageOrderingRulesParsed = true
				continue
			}
			parsePageOrderingRules(line, &pageOrderingRules)
		} else {
			pagesList := make([]int64, 0)
			pages := strings.Split(line, ",")

			for _, page := range pages {
				page, _ := strconv.ParseInt(page, 10, 32)
				pagesList = append(pagesList, page)
			}
			updates = append(updates, pagesList)
		}
	}
	return pageOrderingRules, updates
}
