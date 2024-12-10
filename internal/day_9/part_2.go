package day7

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Range struct {
	begin int
	end   int
	size  int
	value int
}

func SolvePartTwo(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	diskSpace := make([]int, 0)
	emptySpaces := make([]Range, 0)
	fullSpaces := make([]Range, 0)

	for scanner.Scan() {
		diskMap := scanner.Text()

		nFilesProcessed := 0
		index := 0
		for i, fileOrFreeSpace := range diskMap {
			if i%2 == 0 {
				file := fileOrFreeSpace
				fileSize, _ := strconv.ParseInt(string(file), 10, 64)
				currentIndex := index
				for j := 0; j < int(fileSize); j++ {
					diskSpace = append(diskSpace, nFilesProcessed)
					index += 1
				}
				if int(fileSize) > 0 {
					fullSpaces = append(fullSpaces, Range{
						currentIndex, currentIndex + int(fileSize) - 1, int(fileSize), nFilesProcessed,
					})
				}

				nFilesProcessed += 1
			} else {
				freeSpace := fileOrFreeSpace
				spaceSize, _ := strconv.ParseInt(string(freeSpace), 10, 64)
				currentIndex := index
				for j := 0; j < int(spaceSize); j++ {
					diskSpace = append(diskSpace, -1)
					index += 1
				}
				emptySpaces = append(emptySpaces, Range{
					currentIndex, currentIndex + int(spaceSize) - 1, int(spaceSize), -1,
				})
			}
		}
	}

	for i := len(fullSpaces) - 1; i >= 0; i-- {
		fullSpace := fullSpaces[i]

		for j, emptySpace := range emptySpaces {
			if emptySpace.begin > fullSpace.begin {
				break
			}

			if emptySpace.size >= fullSpace.size {
				newSize := emptySpace.size - fullSpace.size

				if newSize == 0 {
					emptySpaces = append(emptySpaces[:j], emptySpaces[j+1:]...)
				} else {
					newBegin := emptySpace.begin + fullSpace.size
					emptySpaces[j].begin = newBegin
					emptySpaces[j].size = newSize
				}

				for k := emptySpace.begin; k < emptySpace.begin+fullSpace.size; k++ {
					diskSpace[k] = fullSpace.value
				}

				for k := fullSpace.begin; k <= fullSpace.end; k++ {
					diskSpace[k] = -1
				}
				break
			}
		}
	}
	result := int64(0)
	for i, value := range diskSpace {
		if value != -1 {
			result += int64(i * value)
			continue
		}
	}

	return fmt.Sprint(result)
}
