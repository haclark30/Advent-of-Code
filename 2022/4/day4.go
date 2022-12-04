package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	input, err := os.ReadFile("input")
	checkErr(err)
	return strings.Split(string(input), "\n")
}

func checkRangesStrict(range1 []int, range2 []int) bool {
	rangeDiff1 := range1[1] - range1[0]
	rangeDiff2 := range2[1] - range2[0]

	if rangeDiff1 < rangeDiff2 {
		if range1[0] >= range2[0] && range1[1] <= range2[1] {
			return true
		}
	} else if rangeDiff2 < rangeDiff1 {
		if range2[0] >= range1[0] && range2[1] <= range1[1] {
			return true
		}
	} else if rangeDiff1 == rangeDiff2 {
		if range1[0] == range2[0] {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func checkRangesOverlap(range1 []int, range2 []int) bool {
	lowerBound := max(range1[0], range2[0])
	upperBound := min(range1[1], range2[1])
	if lowerBound <= upperBound {
		return true
	}
	return lowerBound <= upperBound
}

func rangeStrToInt(rangeStr string) []int {
	rangeInt := make([]int, 2)
	rangeSlice := strings.Split(rangeStr, "-")
	lower, err := strconv.Atoi(rangeSlice[0])
	checkErr(err)
	upper, err := strconv.Atoi(rangeSlice[1])
	checkErr(err)
	rangeInt[0] = lower
	rangeInt[1] = upper
	return rangeInt
}

func part1(pairs []string) int {
	total := 0
	for _, pairStr := range pairs {
		pair := strings.Split(pairStr, ",")
		range1 := rangeStrToInt(pair[0])
		range2 := rangeStrToInt(pair[1])
		if checkRangesStrict(range1, range2) {
			total += 1
		}
	}

	return total
}

func part2(pairs []string) int {
	total := 0
	for _, pairStr := range pairs {
		pair := strings.Split(pairStr, ",")
		range1 := rangeStrToInt(pair[0])
		range2 := rangeStrToInt(pair[1])
		if checkRangesOverlap(range1, range2) {
			total += 1
		}
	}

	return total
}

func main() {
	input := readInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
