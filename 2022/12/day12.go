package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func buildHeightMap(input string) (map[image.Point]rune, image.Point, image.Point) {
	var start, end image.Point
	height := map[image.Point]rune{}
	for x, s := range strings.Fields(string(input)) {
		for y, r := range s {
			height[image.Point{x, y}] = r
			if r == 'S' {
				start = image.Point{x, y}
			} else if r == 'E' {
				end = image.Point{x, y}
			}
		}
	}
	height[start], height[end] = 'a', 'z'

	return height, start, end
}

func bfs(height map[image.Point]rune, start image.Point, end image.Point) (int, int) {
	distances := map[image.Point]int{}
	distances[end] = 0
	bfsQueue := []image.Point{}
	bfsQueue = append(bfsQueue, end)

	var shortest *image.Point

	for len(bfsQueue) > 0 {
		curr := bfsQueue[0]
		bfsQueue = bfsQueue[1:]

		if height[curr] == 'a' && shortest == nil {
			shortest = &curr
		}

		// possible moves as points
		movePoints := []image.Point{
			{0, -1}, {1, 0}, {0, 1}, {-1, 0},
		}

		// make each possible move, if we haven't visited that node yet and it's a valid move
		// add it to the distances map
		for _, d := range movePoints {
			next := curr.Add(d)
			_, seen := distances[next]
			_, valid := height[next]

			if !seen && valid && height[curr] <= height[next]+1 {
				distances[next] = distances[curr] + 1
				bfsQueue = append(bfsQueue, next)
			}
		}
	}

	return distances[start], distances[*shortest]
}

func part1(input string) int {
	height, start, end := buildHeightMap(input)
	pathLen, _ := bfs(height, start, end)
	return pathLen

}

func part2(input string) int {
	height, start, end := buildHeightMap(input)
	_, pathLen := bfs(height, start, end)
	return pathLen
}

func main() {
	input, _ := os.ReadFile("input")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))
}
