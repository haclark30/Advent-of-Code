package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

const add = "addx"
const noop = "noop"

func part1(input []string) int {
	x := 1
	clock := 1
	regSum := 0
	cmdClockMap := map[string]int{
		add:  2,
		noop: 1,
	}
	for _, cmd := range input {
		cmdSplit := strings.Split(cmd, " ")
		cmdClock := cmdClockMap[cmdSplit[0]]
		for i := 0; i < cmdClock; i++ {
			if clock == 20 || (clock-20)%40 == 0 {
				regSum += (x * clock)
			}

			if cmdSplit[0] == add && i == cmdClock-1 {
				// the clock has elapsed, time to actually add
				operand, _ := strconv.Atoi(cmdSplit[1])
				x += operand
			}

			clock++

		}

	}

	return regSum

}

type point struct {
	x int
	y int
}

func inCrtRange(crtPoint point, x int) bool {
	return crtPoint.x >= x-1 && crtPoint.x <= x+1
}

func part2(input []string) {
	x := 1
	clock := 1
	crtPoint := point{x: 0, y: 0}
	output := ""

	cmdClockMap := map[string]int{
		add:  2,
		noop: 1,
	}
	for _, cmd := range input {
		cmdSplit := strings.Split(cmd, " ")
		cmdClock := cmdClockMap[cmdSplit[0]]
		for i := 0; i < cmdClock; i++ {
			if inCrtRange(crtPoint, x) {
				output += "#"
			} else {
				output += "."
			}

			if cmdSplit[0] == add && i == cmdClock-1 {
				// the clock has elapsed, time to actually add
				operand, _ := strconv.Atoi(cmdSplit[1])
				x += operand
			}

			clock++
			crtPoint.x++

			if crtPoint.x == 40 {
				crtPoint.x = 0
				crtPoint.y = 1
				output += "\n"
			}

		}

	}
	fmt.Print(output)

}

func main() {
	input := utils.ReadInput()
	fmt.Println(part1(input))
	part2(input)
}
