package day9

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func SolvePartOne(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	diskSpace := make([]int, 0)

	for scanner.Scan() {
		diskMap := scanner.Text()

		nFilesProcessed := 0
		highestPartition := ((len(diskMap) - 1) / 2) + 1
		reverseIndex := len(diskMap) - 1
		lastFileSizeRemainder := 0
		for i, fileOrFreeSpace := range diskMap {
			if nFilesProcessed == highestPartition {
				for j := 0; j < int(lastFileSizeRemainder); j++ {
					diskSpace = append(diskSpace, nFilesProcessed)
				}
				break
			}
			if i%2 == 0 {
				file := fileOrFreeSpace
				fileSize, _ := strconv.ParseInt(string(file), 10, 32)

				for j := 0; j < int(fileSize); j++ {
					diskSpace = append(diskSpace, nFilesProcessed)
				}
				nFilesProcessed += 1

			} else {
				freeSpace := fileOrFreeSpace
				spaceSize, _ := strconv.ParseInt(string(freeSpace), 10, 32)

				for j := 0; j < int(spaceSize); j++ {
					if lastFileSizeRemainder == 0 {
						rightMostFileSize := diskMap[reverseIndex]
						fileSize, _ := strconv.ParseInt(string(rightMostFileSize), 10, 32)
						lastFileSizeRemainder = int(fileSize)
						highestPartition--
						reverseIndex -= 2
					}

					diskSpace = append(diskSpace, highestPartition)
					lastFileSizeRemainder--
				}
			}
		}
	}

	result := 0
	for i, value := range diskSpace {
		result += i * value
	}

	return fmt.Sprint(result)
}
