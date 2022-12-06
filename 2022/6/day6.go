package main

import (
	"fmt"
	"os"
)

func part1(datastream string) int {
	return findStartOfPattern(datastream, 4)
}

func part2(datastream string) int {
	return findStartOfPattern(datastream, 14)
}

func findStartOfPattern(datastream string, targetLen int) int {
	charmap := make(map[byte]bool)

	for i := 0; i < len(datastream); i++ {
		for j := 0; j <= targetLen-1; j++ {
			char := datastream[i+j]
			if !charmap[char] {
				charmap[char] = true
			} else {
				charmap = make(map[byte]bool)
				break
			}
		}

		if len(charmap) == targetLen {
			return i + targetLen
		}

	}
	return 0
}

func main() {
	input, _ := os.ReadFile("input")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))
}
