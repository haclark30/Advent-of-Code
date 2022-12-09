package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

type Dir struct {
	Size     int
	Path     string
	Children []*Dir
	Parent   *Dir
}

const MAXSIZE = 100000
const DISKSPACE = 70000000
const UNUSED = 30000000

func updateSums(root *Dir, curr *Dir) int {
	if len(curr.Children) == 0 {
		return curr.Size
	}
	for _, child := range curr.Children {
		curr.Size += updateSums(curr, child)
	}
	return curr.Size
}

func getSumRecur(currDir *Dir, currSum int) int {
	if currDir == nil {
		return currSum
	}
	if currDir.Size <= MAXSIZE {
		currSum += currDir.Size
	}

	for _, child := range currDir.Children {
		currSum = getSumRecur(child, currSum)
	}

	return currSum
}

func getSum(curr *Dir) int {
	return getSumRecur(curr, 0)
}

func part1() int {
	sum := 0
	input := utils.ReadInput()
	// going back up to the parent
	// adding a new dir, do i need to check if we've already been here?
	// we're at the root
	// listing contents, continue to the next line
	// we must be listing files, files should update the sum for the current node
	root := createDirTree(input)
	// update sums to include subdirectories

	sum = getSum(&root)

	return sum
}

func findSmalletDir(curr *Dir, space int, currSmallest int) int {
	if curr == nil {
		return currSmallest
	}

	if curr.Size >= space && curr.Size < currSmallest {
		currSmallest = curr.Size
	}

	for _, child := range curr.Children {
		currSmallest = findSmalletDir(child, space, currSmallest)
	}

	return currSmallest

}

func part2() int {
	input := utils.ReadInput()
	root := createDirTree(input)
	totalUnused := DISKSPACE - root.Size
	spaceToFind := UNUSED - totalUnused

	return findSmalletDir(&root, spaceToFind, math.MaxInt)
}

func createDirTree(input []string) Dir {
	currPath := ""
	root := Dir{
		Size: 0,
	}
	currDir := &root

	for _, line := range input {
		lineSplit := strings.Split(line, " ")
		if lineSplit[0] == "$" && lineSplit[1] == "cd" {
			if lineSplit[2] == ".." {

				currDir = currDir.Parent
				currPath = currDir.Path
			} else if lineSplit[2] != "/" {

				currPath += lineSplit[2]
				currPath += "/"
				newDir := Dir{
					Size:   0,
					Path:   currPath,
					Parent: currDir,
				}
				currDir.Children = append(currDir.Children, &newDir)
				currDir = &newDir
			} else {

				currPath += lineSplit[2]
				currDir.Path = currPath
			}

		} else if strings.HasPrefix(line, "$ ls") {

			continue
		} else {

			if lineSplit[0] != "dir" {
				currSize, _ := strconv.Atoi(lineSplit[0])
				currDir.Size += currSize
			}
		}
	}
	root.Size = updateSums(&root, &root)
	return root
}
func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
