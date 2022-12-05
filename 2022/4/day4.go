package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

func checkRangesStrict(range1 []int, range2 []int) bool {
	start1, end1 := range1[0], range1[1]
	start2, end2 := range2[0], range2[1]

	if (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1) {
		return true
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
	return lowerBound <= upperBound

}

func rangeStrToInt(rangeStr string) []int {
	rangeInt := make([]int, 2)
	rangeSlice := strings.Split(rangeStr, "-")
	lower, err := strconv.Atoi(rangeSlice[0])
	utils.CheckErr(err)
	upper, err := strconv.Atoi(rangeSlice[1])
	utils.CheckErr(err)
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
	input := utils.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
