package utils

import (
	"os"
	"strings"
)

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput() []string {
	input, err := os.ReadFile("input")
	CheckErr(err)
	return strings.Split(string(input), "\n")
}
