package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items        []uint64
	op           string // "*" or "+"
	operand      uint64
	testDiv      uint64
	trueMonkey   int
	falseMonkey  int
	inspectCount int
}

func createInitialMonkeys(input string) []*monkey {
	monkeyStrs := strings.Split(input, "\n\n")
	monkeys := make([]*monkey, len(monkeyStrs))

	for i, monkeyStr := range monkeyStrs {
		currMonkeySplit := strings.Split(monkeyStr, "\n")
		itemsSplit := strings.Split(currMonkeySplit[1], ": ")
		opSplit := strings.Split(currMonkeySplit[2], " ")
		testSplit := strings.Split(currMonkeySplit[3], " ")
		testTrueSplit := strings.Split(currMonkeySplit[4], " ")
		testFalseSplit := strings.Split(currMonkeySplit[5], " ")

		currItemsStr := strings.Split(itemsSplit[1], ", ")
		currItems := make([]uint64, len(currItemsStr))
		for j, item := range currItemsStr {
			currItems[j], _ = strconv.ParseUint(item, 10, 64)
		}

		currOp := opSplit[6]
		currOperand, _ := strconv.ParseUint(opSplit[7], 10, 64)
		currTestDiv, _ := strconv.ParseUint(testSplit[len(testSplit)-1], 10, 64)
		currTrueMonkey, _ := strconv.Atoi(testTrueSplit[len(testTrueSplit)-1])
		currFalseMonkey, _ := strconv.Atoi(testFalseSplit[len(testFalseSplit)-1])

		currMonkey := monkey{
			items:        currItems,
			op:           currOp,
			operand:      currOperand,
			testDiv:      currTestDiv,
			trueMonkey:   currTrueMonkey,
			falseMonkey:  currFalseMonkey,
			inspectCount: 0,
		}

		monkeys[i] = &currMonkey
	}
	return monkeys
}

func (m *monkey) turn(monkeys []*monkey, worryLevel uint64, modulo uint64) {
	for _, item := range m.items {
		m.items = m.items[1:]
		var currWorry uint64 = 0
		m.inspectCount++
		if m.op == "+" {
			currWorry = item + m.operand
		} else {
			currWorry = item * m.operand
		}

		if m.operand == 0 && m.op == "+" {
			currWorry = item + item
		} else if m.operand == 0 && m.op == "*" {
			currWorry = item * item
		}

		if currWorry >= math.MaxUint64-1 {
			fmt.Println(currWorry)
		}

		currWorry /= worryLevel
		if worryLevel == 1 {
			currWorry %= modulo

		}

		if currWorry%m.testDiv == 0 {
			monkeys[m.trueMonkey].items = append(monkeys[m.trueMonkey].items, currWorry)
		} else {
			monkeys[m.falseMonkey].items = append(monkeys[m.falseMonkey].items, currWorry)
		}

	}

}

func calcMonkeyBusiness(monkeys []*monkey) int {
	large1 := 0
	large2 := 0

	large1 = monkeys[0].inspectCount
	for i := 1; i < len(monkeys); i++ {
		if large1 < monkeys[i].inspectCount {
			large2 = large1
			large1 = monkeys[i].inspectCount
		} else if large2 < monkeys[i].inspectCount {
			large2 = monkeys[i].inspectCount
		}
	}
	return large1 * large2
}

func part1(input string) int {
	monkeys := createInitialMonkeys(input)
	numrounds := 20
	for i := 0; i < numrounds; i++ {
		for _, monkey := range monkeys {
			monkey.turn(monkeys, 3, 0)
		}
	}

	return calcMonkeyBusiness(monkeys)
}

func part2(input string) int {
	monkeys := createInitialMonkeys(input)
	numrounds := 10000

	// modulo black magic
	modulo := 1
	for _, monkey := range monkeys {
		modulo *= int(monkey.testDiv)
	}

	for i := 0; i < numrounds; i++ {
		for _, monkey := range monkeys {
			monkey.turn(monkeys, 1, uint64(modulo))
		}
	}

	return calcMonkeyBusiness(monkeys)
}
func main() {
	input, _ := os.ReadFile("input")
	fmt.Println(part1(string(input)))
	fmt.Println(part2(string(input)))
}
