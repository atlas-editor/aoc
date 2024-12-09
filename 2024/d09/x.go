package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]

	data, _ := os.ReadFile(path)
	input := strings.TrimSpace(string(data))

	fmt.Println(p1(input))
	fmt.Println(p2(input))
}

func p1(input string) int {
	diskMap := parse(input)

	freeSpacePtr := 0
	fileBlockPtr := len(diskMap) - 1
	for {
		for fileBlockPtr >= 0 && diskMap[fileBlockPtr] == -1 {
			fileBlockPtr--
		}
		for freeSpacePtr < len(diskMap) && diskMap[freeSpacePtr] != -1 {
			freeSpacePtr++
		}
		if freeSpacePtr >= fileBlockPtr {
			break
		}
		diskMap[freeSpacePtr], diskMap[fileBlockPtr] = diskMap[fileBlockPtr], -1
	}

	s := 0
	for k := range diskMap {
		if diskMap[k] == -1 {
			continue
		}
		s += k * diskMap[k]
	}

	return s
}

func p2(input string) int {
	diskMap := parse(input)

	var upperLimit, freeSpacePtr, fileBlockPtr, fileId, fileBlockLen int
	var found bool
	upperLimit = len(diskMap) - 1
	for {
		fileBlockPtr, upperLimit, fileId, found = findPrevFileBlock(diskMap, upperLimit)
		if !found {
			break
		}

		fileBlockLen = upperLimit - fileBlockPtr

		freeSpacePtr, found = findFreeSpaceBlock(diskMap, fileBlockLen, fileBlockPtr)

		if found {
			for k := range fileBlockLen {
				diskMap[fileBlockPtr+k] = -1
				diskMap[freeSpacePtr+k] = fileId
			}
		}

		upperLimit = fileBlockPtr - 1
	}

	s := 0
	for k := range diskMap {
		if diskMap[k] == -1 {
			continue
		}
		s += k * diskMap[k]
	}

	return s
}

func parse(input string) []int {
	diskMap := []int{}
	for i := range input {
		for range input[i] - 48 {
			if i%2 == 0 {
				diskMap = append(diskMap, i/2)
			} else {
				diskMap = append(diskMap, -1)
			}
		}
	}
	return diskMap
}

func findPrevFileBlock(diskMap []int, limit int) (int, int, int, bool) {
	for limit >= 0 && diskMap[limit] == -1 {
		limit--
	}
	if limit < 0 {
		return -1, -1, -1, false
	}
	curr := diskMap[limit]
	j1 := limit
	for j1 >= 0 && diskMap[j1] == curr {
		j1--
	}
	return j1 + 1, limit + 1, curr, true
}

func findFreeSpaceBlock(diskMap []int, size, limit int) (int, bool) {
	i, i1 := 0, 0
	for i1-i < size && i < limit {
		for i < limit && diskMap[i] != -1 {
			i++
		}
		if i >= limit {
			break
		}
		i1 = i
		for i1 < limit && diskMap[i1] == -1 {
			i1++
		}
		if i1-i >= size {
			return i, true
		}
		i = i1
	}

	return -1, false
}
