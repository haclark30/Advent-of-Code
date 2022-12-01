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

func findMax(elfs []string) int {
	max := 0
	for i := 0; i < len(elfs); i++ {
		elfCals := strings.Split(elfs[i], "\n")
		elfCalTotal := 0
		for j := 0; j < len(elfCals); j++ {
			intCal, err := strconv.Atoi(elfCals[j])
			checkErr(err)
			elfCalTotal += intCal
		}
		if elfCalTotal > max {
			max = elfCalTotal
		}
	}
	return max
}

func findTopThreeTotal(elfs []string) int {
	max1 := 0
	max2 := -1
	max3 := -2
	for i := 0; i < len(elfs); i++ {
		elfCals := strings.Split(elfs[i], "\n")
		elfCalTotal := 0
		for j := 0; j < len(elfCals); j++ {
			intCal, err := strconv.Atoi(elfCals[j])
			checkErr(err)
			elfCalTotal += intCal
		}
		if elfCalTotal > max1 {
			tempMax1 := max1
			tempMax2 := max2
			max1 = elfCalTotal
			// need to shift the previous highest down
			max2 = tempMax1
			max3 = tempMax2

		} else if elfCalTotal > max2 {
			tempMax2 := max2
			max2 = elfCalTotal
			max3 = tempMax2
		} else if elfCalTotal > max3 {
			max3 = elfCalTotal
		}
	}
	return max1 + max2 + max3
}

func main() {
	input, err := os.ReadFile("input")
	checkErr(err)
	inputStr := string(input)
	elfSplit := strings.Split(inputStr, "\n\n")

	max := findMax(elfSplit)
	topThree := findTopThreeTotal(elfSplit)
	fmt.Println(max)
	fmt.Println(topThree)
}
