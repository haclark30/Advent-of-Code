package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

type point struct {
	x int
	y int
}

func (p *point) move(dir string) {
	switch dir {
	case "R":
		p.x++
	case "U":
		p.y++
	case "L":
		p.x--
	case "D":
		p.y--
	}
}

func checkAdjacency(p1 point, p2 point) bool {
	xDiff := p1.x - p2.x
	yDiff := p1.y - p2.y
	xDiff = int(math.Abs(float64(xDiff)))
	yDiff = int(math.Abs(float64(yDiff)))

	return xDiff <= 1 && yDiff <= 1
}

func makeTailMove(head *point, tail *point) {
	if head.x == tail.x {
		//same column
		if head.y > tail.y {
			tail.move("U")
		} else {
			tail.move("D")
		}
	} else if head.y == tail.y {
		// same row
		if head.x > tail.x {
			tail.move("R")
		} else {
			tail.move("L")
		}
	} else {
		// different row or col, need to move diagonally
		// check y
		if head.y > tail.y {
			tail.move("U")
		} else {
			tail.move("D")
		}
		// check x
		if head.x > tail.x {
			tail.move("R")
		} else {
			tail.move("L")
		}
	}
}

func part1(input []string) int {
	head := point{x: 0, y: 0}
	tail := point{x: 0, y: 0}
	pointMap := make(map[string]bool)

	for _, move := range input {
		moveSplit := strings.Split(move, " ")
		dir := moveSplit[0]
		num, _ := strconv.Atoi(moveSplit[1])

		for i := 0; i < num; i++ {
			head.move(dir)
			if checkAdjacency(head, tail) {
				// already ajacent, move on
				posString := fmt.Sprintf("%d,%d", tail.x, tail.y)
				pointMap[posString] = true
				continue
			}
			makeTailMove(&head, &tail)
			posString := fmt.Sprintf("%d,%d", tail.x, tail.y)
			pointMap[posString] = true
		}
	}

	return len(pointMap)
}

func part2(input []string) int {
	// todo: part1 and part2 can be combined, don't need head vs tail, just keep list of points
	head := point{x: 0, y: 0}
	var tails [9]point
	for i := 0; i < 9; i++ {
		tails[i] = point{x: 0, y: 0}
	}
	pointMap := make(map[string]bool)

	// similar to part 1
	for _, move := range input {
		moveSplit := strings.Split(move, " ")
		dir := moveSplit[0]
		num, _ := strconv.Atoi(moveSplit[1])

		for i := 0; i < num; i++ {
			// head checks against the first tail
			head.move(dir)
			if checkAdjacency(head, tails[0]) {
				continue
			}
			// move all of the tails relative to the one in front of it
			makeTailMove(&head, &tails[0])
			for j := 1; j < len(tails); j++ {
				if !checkAdjacency(tails[j-1], tails[j]) {
					makeTailMove(&tails[j-1], &tails[j])
				}

			}
			posString := fmt.Sprintf("%d,%d", tails[8].x, tails[8].y)
			pointMap[posString] = true
		}
	}
	return len(pointMap)
}

func main() {
	input := utils.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
