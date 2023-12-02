package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/haclark30/Advent-of-Code/utils"
)
type Round struct {
	Red int
	Blue int
	Green int
}

var MaxRound = Round{
	Red: 12,
	Green: 13,
	Blue: 14,

}

type Game struct {
	Rounds []Round
}

func main() {
	part1()
	part2()
}

func part1() {

	input := utils.ReadInput()
	sum := 0
	for idx , val := range input {
		gameId := idx + 1
		gameStr := strings.Split(val, ": ")[1]
		rounds := strings.Split(gameStr, "; ")
		var game Game
		for _, r := range rounds {
			round := NewRound(r)
			game.Rounds = append(game.Rounds, round)
		}
		gamePossbile := IsGamePossible(game)
		if gamePossbile {
			sum += gameId
		}
	}
	fmt.Println(sum)
}

func part2() {
	input := utils.ReadInput()
	sum := 0
	for _, val := range input {
		gameStr := strings.Split(val, ": ")[1]
		rounds := strings.Split(gameStr, "; ")
		var game Game
		for _, r := range rounds {
			round := NewRound(r)
			game.Rounds = append(game.Rounds, round)
		}
		minRound := MinRound(game)
		power := minRound.Red * minRound.Green * minRound.Blue
		sum += power
	}
	fmt.Println(sum)
}

func NewRound(roundStr string) Round {
	round := Round{}
	split := strings.Split(roundStr, ", ")
	for _, s := range split {
		colorSplit := strings.Split(s, " ")
		colorVal, _ := strconv.Atoi(colorSplit[0])
		switch colorSplit[1] {
		case "red":
			round.Red = colorVal
		case "blue":
			round.Blue = colorVal
		case "green":
			round.Green = colorVal
		}
	}
	return round
}

func IsGamePossible(game Game) bool {
	for _, r := range game.Rounds {
		red := MaxRound.Red - r.Red
		green := MaxRound.Green - r.Green
		blue := MaxRound.Blue - r.Blue
		if red < 0 || green < 0 || blue < 0 {
			return false
		}
	}
	
	return true
}

func MinRound(game Game) Round {
	var minRound Round
	for _, r := range game.Rounds {
		if r.Red > minRound.Red {
			minRound.Red = r.Red
		}
		if r.Green > minRound.Green {
			minRound.Green = r.Green
		}

		if r.Blue > minRound.Blue {
			minRound.Blue = r.Blue
		}
	}
	return minRound
}
