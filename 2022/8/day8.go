package main

import (
	"fmt"

	"github.com/haclark30/Advent-of-Code/utils"
)

func buildGrid(input []string) [][]int {
	grid := make([][]int, len(input))
	for i := range grid {
		grid[i] = make([]int, len(input[0]))
	}

	for i := range input {
		for j, val := range input[i] {
			grid[i][j] = int(val - '0')
		}
	}
	return grid
}

func countVisible(visibleGrid [][]bool) int {
	visible := 0
	for _, row := range visibleGrid {
		for _, col := range row {
			if col {
				visible++
			}
		}
	}
	return visible
}

func part1(grid [][]int) int {
	visibleGrid := make([][]bool, len(grid))
	for i := range visibleGrid {
		visibleGrid[i] = make([]bool, len(grid[0]))
	}

	for i := range grid {
		for j := range grid[i] {
			// all outer trees are visible
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[0])-1 {
				visibleGrid[i][j] = true
				continue
			}

			treeVal := grid[i][j]
			visibleGrid[i][j] = true

			upVis := true
			downVis := true
			leftVis := true
			rightVis := true

			//check up
			for idx := i - 1; idx >= 0; idx-- {
				if grid[idx][j] >= treeVal {
					upVis = false
				}
			}
			// check down
			for idx := i + 1; idx < len(grid); idx++ {
				if grid[idx][j] >= treeVal {
					downVis = false
				}
			}

			// check left
			for idx := j - 1; idx >= 0; idx-- {
				if grid[i][idx] >= treeVal {
					leftVis = false
				}
			}

			//check right
			for idx := j + 1; idx < len(grid[0]); idx++ {
				if grid[i][idx] >= treeVal {
					rightVis = false
				}
			}

			// if visible from any direction, it is visilbe
			visibleGrid[i][j] = upVis || downVis || leftVis || rightVis

		}
	}

	return countVisible(visibleGrid)
}

func part2(grid [][]int) int {
	bestScore := 0
	for row := range grid {
		for col := range grid {
			upScore, downScore, leftScore, rightScore := 0, 0, 0, 0
			treeVal := grid[row][col]
			//check up
			for idx := row - 1; idx >= 0; idx-- {
				upScore += 1
				if grid[idx][col] >= treeVal {
					break
				}
			}

			//check down
			for idx := row + 1; idx < len(grid); idx++ {
				downScore += 1
				if grid[idx][col] >= treeVal {
					break
				}
			}

			// check left
			for idx := col - 1; idx >= 0; idx-- {
				leftScore += 1
				if grid[row][idx] >= treeVal {
					break
				}
			}

			//check right
			for idx := col + 1; idx < len(grid[0]); idx++ {
				rightScore += 1
				if grid[row][idx] >= treeVal {
					break
				}
			}

			currScore := upScore * downScore * rightScore * leftScore
			if currScore > bestScore {
				bestScore = currScore
			}
		}
	}
	return bestScore
}

func main() {
	input := utils.ReadInput()
	grid := buildGrid(input)
	fmt.Println(part1(grid))
	grid = buildGrid(input)
	fmt.Println(part2(grid))
}
