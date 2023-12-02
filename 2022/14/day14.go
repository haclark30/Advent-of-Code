package main

import (
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

func buildPointMap(input []string) map[image.Point]bool {
	pointMap := map[image.Point]bool{}

	for _, ln := range input {
		ln = strings.ReplaceAll(ln, " ", "")
		pts := strings.Split(ln, "->")
		for i := 0; i+1 < len(pts); i++ {
			xy := strings.Split(pts[i], ",")
			currX, _ := strconv.Atoi(xy[0])
			currY, _ := strconv.Atoi(xy[1])
			currPt := image.Point{currX, currY}
			xy = strings.Split(pts[i+1], ",")
			currX, _ = strconv.Atoi(xy[0])
			currY, _ = strconv.Atoi(xy[1])
			nextPt := image.Point{currX, currY}

			var addPt image.Point
			subPt := currPt.Sub(nextPt)
			if subPt.X == 0 && subPt.Y < 0 {
				addPt = image.Point{0, 1}
			} else if subPt.X == 0 && subPt.Y > 0 {
				addPt = image.Point{0, -1}
			} else if subPt.X > 0 && subPt.Y == 0 {
				addPt = image.Point{-1, 0}
			} else {
				addPt = image.Point{1, 0}
			}

			for !currPt.Eq(nextPt) {
				pointMap[currPt] = true
				currPt = currPt.Add(addPt)
			}
			pointMap[currPt] = true

		}
	}
	return pointMap
}

func part1(input []string) int {
	pointMap := buildPointMap(input)
	lowestPt := image.Point{500, 0}
	for pt := range pointMap {
		if pt.Y > lowestPt.Y {
			lowestPt = pt
		}
	}
	source := image.Point{500, 0}
	currSand := source
	numSand := 0
	// keep adding sand until we are past the lowest point in the structures
	for currSand.Y < lowestPt.Y {

		// check down
		if p := currSand.Add(image.Point{0, 1}); !pointMap[p] {
			currSand = p
			continue
		}
		// check down-left
		if p := currSand.Add(image.Point{-1, 1}); !pointMap[p] {
			currSand = p
			continue
		}

		// check down-right
		if p := currSand.Add(image.Point{1, 1}); !pointMap[p] {
			currSand = p
			continue
		}

		//no where to go, add to blocked
		pointMap[currSand] = true
		currSand = source
		numSand++

	}
	return numSand
}

func part2(input []string) int {
	pointMap := buildPointMap(input)
	lowestPt := image.Point{500, 0}
	for pt := range pointMap {
		if pt.Y > lowestPt.Y {
			lowestPt = pt
		}
	}

	floor := lowestPt.Y + 2
	source := image.Point{500, 0}
	currSand := source
	numSand := 0
	// keep adding sand until we are past the lowest point in the structures
	for {
		// check down
		if p := currSand.Add(image.Point{0, 1}); !pointMap[p] && p.Y < floor {
			currSand = p
			continue
		}
		// check down-left
		if p := currSand.Add(image.Point{-1, 1}); !pointMap[p] && p.Y < floor {
			currSand = p
			continue
		}

		// check down-right
		if p := currSand.Add(image.Point{1, 1}); !pointMap[p] && p.Y < floor {
			currSand = p
			continue
		}

		//no where to go, add to blocked
		pointMap[currSand] = true
		numSand++
		if currSand.Eq(source) {
			break
		}
		currSand = source

	}
	return numSand
}

func main() {
	input := utils.ReadInput()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
