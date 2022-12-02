package main

import (
	"fmt"
	"os"
	"strings"
)

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

func scoreGame(player1 string, player2 string, scoreMap map[string]int, winLoseMap map[string]map[string]int) int {
	playerScore := scoreMap[player1]
	resultScore := 0
	// ties return player score plus 3
	if player1 == player2 {
		resultScore = 3
		return playerScore + resultScore
	}

	return playerScore + winLoseMap[player1][player2]
}

func part1(input string) int {
	totalScore := 0
	scoreMap := makeScoreMap()

	// A == Rock, B == Paper, C == Scissors
	// X == Rock, Y == Paper, Z == Scissors
	// turn player2 into same scheme as player 1 to make it easier
	moveMap := makeMoveMap()
	winLoseMap := makeWinLoseMap()

	gameList := strings.Split(input, "\n")
	for i := 0; i < len(gameList); i++ {
		gameSplit := strings.Split(gameList[i], " ")
		player1 := moveMap[gameSplit[1]]
		player2 := moveMap[gameSplit[0]]
		totalScore += scoreGame(player1, player2, scoreMap, winLoseMap)
	}
	return totalScore
}

func makeWinLoseMap() map[string]map[string]int {
	winLoseMap := make(map[string]map[string]int)
	winLoseMap["ROCK"] = make(map[string]int)
	winLoseMap["PAPER"] = make(map[string]int)
	winLoseMap["SCISSORS"] = make(map[string]int)

	winLoseMap["ROCK"]["PAPER"] = 0
	winLoseMap["ROCK"]["SCISSORS"] = 6
	winLoseMap["PAPER"]["ROCK"] = 6
	winLoseMap["PAPER"]["SCISSORS"] = 0
	winLoseMap["SCISSORS"]["ROCK"] = 0
	winLoseMap["SCISSORS"]["PAPER"] = 6
	return winLoseMap
}

func makeMoveMap() map[string]string {
	moveMap := make(map[string]string)
	moveMap["A"] = "ROCK"
	moveMap["B"] = "PAPER"
	moveMap["C"] = "SCISSORS"

	moveMap["X"] = "ROCK"
	moveMap["Y"] = "PAPER"
	moveMap["Z"] = "SCISSORS"
	return moveMap
}

func findIdealMove(opponentMove string, outcome string) string {
	// X = LOSE, Y = DRAW, Z = WIN
	move := ""
	if outcome == "Y" {
		move = opponentMove
	}
	if outcome == "X" {
		switch opponentMove {
		case "ROCK":
			move = "SCISSORS"
		case "PAPER":
			move = "ROCK"
		case "SCISSORS":
			move = "PAPER"
		}
	} else if outcome == "Z" {
		switch opponentMove {
		case "ROCK":
			move = "PAPER"
		case "PAPER":
			move = "SCISSORS"
		case "SCISSORS":
			move = "ROCK"
		}
	}
	return move
}

func makeScoreMap() map[string]int {
	scoreMap := make(map[string]int)
	scoreMap["ROCK"] = 1
	scoreMap["PAPER"] = 2
	scoreMap["SCISSORS"] = 3
	return scoreMap
}

func part2(input string) int {
	totalScore := 0
	moveMap := makeMoveMap()
	scoreMap := makeScoreMap()
	winLoseMap := makeWinLoseMap()

	gameList := strings.Split(input, "\n")
	for i := 0; i < len(gameList); i++ {
		gameSplit := strings.Split(gameList[i], " ")
		player2 := moveMap[gameSplit[0]]
		player1 := findIdealMove(player2, gameSplit[1])
		totalScore += scoreGame(player1, player2, scoreMap, winLoseMap)
	}
	return totalScore
}

func main() {
	input := readInput()
	part1_ans := part1(input)
	part2_ans := part2(input)
	fmt.Println(part1_ans)
	fmt.Println(part2_ans)
}
