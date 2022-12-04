package main

import (
	"os"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() []string {
	input, err := os.ReadFile("input")
	checkErr(err)
	return strings.Split(string(input), "\n")
}

func part1() {

}

func main() {
	input := readInput()
}
