package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func manhattan(p1 image.Point, p2 image.Point) int {
	return int(math.Abs(
		float64(p1.X)-float64(p2.X)) + math.Abs(
		float64(p1.Y)-float64(p2.Y)))
}

func part1(input []string) {
	f := func(c rune) bool {
		return !unicode.IsNumber(c) && c != '-'
	}
	cave := map[image.Point]bool{}
	target := 10
	sum := 0

	for _, str := range input {
		coords := strings.FieldsFunc(str, f)
		srcX, _ := strconv.Atoi(coords[0])
		srcY, _ := strconv.Atoi(coords[1])
		dstX, _ := strconv.Atoi(coords[2])
		dstY, _ := strconv.Atoi(coords[3])
		source := image.Point{srcX, srcY}
		dest := image.Point{dstX, dstY}

		dist := manhattan(source, dest)
		rowsAway := source.Y - target

		if int(math.Abs(float64(rowsAway))) > dist {
			continue
		}

		numPossible := (dist*2 + 1) - 2*int(math.Abs(float64(rowsAway)))
		sum += numPossible
		for i := -(numPossible - 1) / 2; i < (numPossible-1)/2; i++ {
			cave[image.Point{i + source.X, target}] = true
		}

	}
	fmt.Println(len(cave))

}

func part2(input []string) {
	f := func(c rune) bool {
		return !unicode.IsNumber(c) && c != '-'
	}
	cave := map[image.Point]int{}
	tuning := 4000000
	size := 4000000

	for _, str := range input {
		coords := strings.FieldsFunc(str, f)
		srcX, _ := strconv.Atoi(coords[0])
		srcY, _ := strconv.Atoi(coords[1])
		dstX, _ := strconv.Atoi(coords[2])
		dstY, _ := strconv.Atoi(coords[3])
		source := image.Point{srcX, srcY}
		dest := image.Point{dstX, dstY}
		dist := manhattan(source, dest)
		cave[source] = dist
	}

	for y := 0; y <= size; y++ {
		for x := 0; x <= size; x++ {
			inside := false
			for s, dist := range cave {
				if manhattan(s, image.Point{x, y}) <= dist {
					inside = true
					x = s.X + dist - int(math.Abs(float64(s.Y-y)))
					break
				}
			}
			if !inside {
				fmt.Println(x*tuning + y)
			}
		}
	}
}

func main() {
	input, _ := os.ReadFile("input")
	part1(strings.Split(string(input), "\n"))
	part2(strings.Split(string(input), "\n"))
}
