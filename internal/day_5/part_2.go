package day5

import (
	"bufio"
	"fmt"
	"strings"
)

type Page struct {
	index int64
	value int64
}

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	pageOrderingRules, updates := parsePuzzleInput(scanner)
	result := int64(0)

	for _, update := range updates {
		result += getMiddleValueForWronglyOrderedRows(pageOrderingRules, update)
	}

	return fmt.Sprint(result)
}

func getMiddleValueForWronglyOrderedRows(pageOrderingRules map[int64]Set, updateRow []int64) int64 {
	requirementsForSeenPages := make(Set)
	seenPages := make([]int64, 0)
	invalidlyOrdered := false

	for _, page := range updateRow {
		pageStillAllowed, _ := isPageStillAllowed(page, requirementsForSeenPages)
		if !pageStillAllowed {
			invalidlyOrdered = true
			shouldBeBeforePages := pageOrderingRules[page]

			highestIndex := 0
			for beforePage := range shouldBeBeforePages {
				for i, seenPage := range seenPages {
					if beforePage == seenPage {
						if highestIndex < i+1 {
							highestIndex = i + 1
						}
						break
					}
				}
				requirementsForSeenPages[beforePage] = struct{}{}
			}
			if highestIndex == 0 {
				left := make([]int64, 0)
				left = append(left, page)
				right := seenPages
				seenPages = append(left, right...)
			} else {
				newUpdateSlice := make([]int64, 0)
				newUpdateSlice = append(newUpdateSlice, seenPages[:highestIndex]...)
				newUpdateSlice = append(newUpdateSlice, page)
				newUpdateSlice = append(newUpdateSlice, seenPages[highestIndex:]...)
				seenPages = newUpdateSlice
			}
		} else {
			shouldBeBeforePages := pageOrderingRules[page]

			for shouldBeBeforePage := range shouldBeBeforePages {
				requirementsForSeenPages[shouldBeBeforePage] = struct{}{}
			}
			seenPages = append(seenPages, page)
		}
	}

	result := int64(0)
	if invalidlyOrdered {
		result = seenPages[(len(seenPages)-1)/2]
	}
	return result
}
