package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

var numMap = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	input := utils.ReadInput()
	sum := 0
	for _, s := range input {
		sum += getCalibrationValue(s)
	}
	fmt.Println(sum)
}

func getCalibrationValue(s string) int {
	for k, v := range numMap {
		s = strings.Replace(s, k, v, -1)
	}
	calibrationVals := make([]int, 0)
	for _, c := range s {
		val, err := strconv.Atoi(string(c))
		if err != nil {
			continue
		}
		calibrationVals = append(calibrationVals, val)
	}
	if len(calibrationVals) > 0 {
		return (calibrationVals[0] * 10) + calibrationVals[len(calibrationVals)-1]
	}
	return 0
}
