package main

import (
	"fmt"

	"github.com/haclark30/Advent-of-Code/utils"
)

func letterToValue(c byte) int {
	intVal := int(c)
	if intVal >= 96 {
		return intVal - 96
	}
	return intVal - 38
}

func splitStringHalf(s string) (string, string) {
	halfPoint := len(s) / 2
	s1 := s[0:halfPoint]
	s2 := s[halfPoint:]
	return s1, s2
}

func findCommonLetter(rucksack string) byte {
	comp1, comp2 := splitStringHalf(rucksack)
	comp1Map := make(map[byte]bool)

	for i := 0; i < len(comp1); i++ {
		comp1Map[comp1[i]] = true
	}

	for i := 0; i < len(comp2); i++ {
		if comp1Map[comp2[i]] {
			return comp2[i]
		}
	}
	fmt.Println(comp1, comp2)
	return 'A'
}

func findCommonLetterGroup(rucksacks []string) byte {
	letterMap := make(map[byte]bool)
	for i := 0; i < len(rucksacks[0]); i++ {
		letterMap[rucksacks[0][i]] = true
	}

	for i := 1; i < len(rucksacks); i++ {
		nextRucksack := rucksacks[i]
		candidates := make(map[byte]bool)
		for j := 0; j < len(nextRucksack); j++ {
			if letterMap[nextRucksack[j]] {
				candidates[nextRucksack[j]] = true
			}
		}

		for key := range letterMap {
			if !candidates[key] {
				delete(letterMap, key)
			}
		}
	}

	for key := range letterMap {
		return key
	}
	return 'A'
}

func part1(input []string) int {
	total := 0
	for i := 0; i < len(input); i++ {
		total += letterToValue(findCommonLetter(input[i]))
	}

	return total
}

func part2(input []string) int {
	total := 0
	for i := 0; i < len(input); i += 3 {
		total += letterToValue(findCommonLetterGroup(input[i : i+3]))
	}
	return total
}

func main() {
	input := utils.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
