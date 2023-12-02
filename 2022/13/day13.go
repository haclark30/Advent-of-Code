package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)

func part1(pairs []string) {
	pairIndex := 1
	sum := 0
	var all []any
	for i := 0; i+1 < len(pairs); i += 2 {
		left := pairs[i]
		right := pairs[i+1]

		var first any
		var second any

		err := json.Unmarshal([]byte(left), &first)
		utils.CheckErr(err)
		err = json.Unmarshal([]byte(right), &second)
		utils.CheckErr(err)
		all = append(all, first, second)

		if comparePackets(first, second) <= 0 {
			sum += pairIndex
		}
		pairIndex++
	}

	fmt.Println(sum)

	var add1 any
	var add2 any
	json.Unmarshal([]byte("[[2]]"), &add1)
	json.Unmarshal([]byte("[[6]]"), &add2)
	all = append(all, add1, add2)

	sort.Slice(all, func(i, j int) bool {
		return comparePackets(all[i], all[j]) < 0
	})

	var r int = 1
	for k, v := range all {
		str, _ := json.Marshal(v)
		if string(str) == "[[2]]" || string(str) == "[[6]]" {
			r *= k + 1
		}
	}

	fmt.Println(r)

}

func comparePackets(left any, right any) int {
	l, lOK := left.(float64)
	r, rOK := right.(float64)
	if lOK && rOK {
		return int(l) - int(r)
	}

	var lList []any
	var rList []any

	switch left.(type) {
	case []any, []float64:
		lList = left.([]any)
	case float64:
		lList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		rList = right.([]any)
	case float64:
		rList = []any{right}
	}

	for i := range lList {
		if len(rList) <= i {
			return 1
		}
		if r := comparePackets(lList[i], rList[i]); r != 0 {
			return r
		}
	}

	if len(lList) == len(rList) {
		return 0
	}

	return -1
}

func main() {
	input, _ := os.ReadFile("input")
	part1(strings.Fields(string(input)))
}
