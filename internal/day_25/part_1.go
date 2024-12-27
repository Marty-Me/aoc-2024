package day_25

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Locks = [][]string
type Keys = [][]string

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	locks, keys := parseLockAndKeys(scanner)
	nKeysFit := 0
	for _, lock := range locks {
		for _, key := range keys {
			teethFit := 0
			for i := 0; i < 5; i++ {
				lockInt, _ := strconv.ParseInt(lock[i], 2, 16)
				keyInt, _ := strconv.ParseInt(key[i], 2, 16)

				if (lockInt & keyInt) == 0 {
					teethFit++
				} else {
					break
				}
			}
			if teethFit == 5 {
				nKeysFit++
			}
		}
	}

	return fmt.Sprint(nKeysFit)
}

func parseLockAndKeys(scanner *bufio.Scanner) (Locks, Keys) {
	keys := make(Locks, 0)
	locks := make(Keys, 0)

	keyPart := make([]string, 5)

	y := 0
	isKey := false

	for scanner.Scan() {
		line := scanner.Text()

		if (y % 8) == 0 {
			isKey = line[0] != '#'
		}

		// Read each pin of the lock/key and store it in binary format
		for i, char := range line {
			if char == '#' {
				newBit := "1" + strings.Repeat("0", 6-y%8)
				newInt, _ := strconv.ParseInt(newBit, 2, 16)

				existing := keyPart[i]
				existingInt, _ := strconv.ParseInt(existing, 2, 16)

				intStr := strconv.FormatInt(newInt|existingInt, 2)
				keyPart[i] = strings.Repeat("0", 6-(len(intStr)-1)) + intStr
			}
		}

		// When the last pin of the lock/key has been processed, sort it in the keys or locks slice
		if y != 0 && ((y % 8) == 6) {
			if isKey {
				keys = append(keys, keyPart)
			} else {
				locks = append(locks, keyPart)
			}
			keyPart = make([]string, 5)
		}
		y++
	}
	return locks, keys
}
