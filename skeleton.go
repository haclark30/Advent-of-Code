package main

import "os"

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput() string {
	input, err := os.ReadFile("input")
	checkErr(err)
	return string(input)
}

func part1() {

}

func main() {
	input := readInput()
}
