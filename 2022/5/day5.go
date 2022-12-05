package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

func buildStacks(input []string) [][]byte {
	var rows [][]byte
	var stacks [][]byte

	// turn the input into a list of lists of bytes
	for _, row := range input {
		if row[0:2] == " 1" {
			break
		}
		var byteRow []byte
		for i := 1; i < len(row); i += 4 {
			byteRow = append(byteRow, row[i : i+1][0])
		}
		rows = append(rows, byteRow)
	}

	// turn the rows into stacks
	for i := 0; i < len(rows[0]); i++ {
		var stack []byte
		for j := len(rows) - 1; j >= 0; j-- {
			if rows[j][i] != ' ' {
				stack = append(stack, rows[j][i])
			}
		}
		stacks = append(stacks, stack)
	}

	return stacks
}

func createMoves(input []string) [][]int {
	var moves [][]int
	for _, row := range input {
		if strings.HasPrefix(row, "move") {
			rowSplit := strings.Split(row, " ")
			move1, err := strconv.Atoi(rowSplit[1])
			utils.CheckErr(err)
			move2, err := strconv.Atoi(rowSplit[3])
			utils.CheckErr(err)
			move3, err := strconv.Atoi(rowSplit[5])
			utils.CheckErr(err)
			move := []int{move1, move2, move3}
			moves = append(moves, move)
		}
	}
	return moves
}

func part1(stacks [][]byte, moves [][]int) {
	for _, move := range moves {
		moveLen := move[0]
		fromIdx := move[1] - 1
		toIdx := move[2] - 1

		for i := 0; i < moveLen; i++ {
			itemToMove := stacks[fromIdx][len(stacks[fromIdx])-1]
			stacks[toIdx] = append(stacks[toIdx], itemToMove)
			stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-1]
		}
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Println("")
}

func part2(stacks [][]byte, moves [][]int) {
	for _, move := range moves {
		moveLen := move[0]
		fromIdx := move[1] - 1
		toIdx := move[2] - 1

		var newMove []byte

		for i := 0; i < moveLen; i++ {
			itemToMove := stacks[fromIdx][len(stacks[fromIdx])-1]
			newMove = append(newMove, itemToMove)
			stacks[fromIdx] = stacks[fromIdx][:len(stacks[fromIdx])-1]
		}

		for i := len(newMove) - 1; i >= 0; i-- {
			moveItem := newMove[i]
			stacks[toIdx] = append(stacks[toIdx], moveItem)
		}
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Println("")
}

func main() {
	input := utils.ReadInput()
	stacks := buildStacks(input)
	moves := createMoves(input)
	part1(stacks, moves)
	stacks = buildStacks(input)
	moves = createMoves(input)
	part2(stacks, moves)
}
