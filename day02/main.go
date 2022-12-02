package main

import (
	"fmt"
	"strings"

	"github.com/justingodden/advent-of-code-2022/utils"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	input := utils.ReadFile("input.txt")
	rounds := strings.Split(input, "\n")
	var playerScore int

	for _, round := range rounds {
		roundVals := strings.Split(round, " ")
		cpu := roundVals[0]
		player := roundVals[1]
		playerScore += shapeScore[player]
		playerScore += roundScore[player][cpu]
	}
	return playerScore
}

func part2() int {
	input := utils.ReadFile("input.txt")
	rounds := strings.Split(input, "\n")
	var playerScore int

	for _, round := range rounds {
		roundVals := strings.Split(round, " ")
		cpu := roundVals[0]
		outcome := roundVals[1]
		playerScore += outcomeScore[outcome]
		playerScore += outcomeToChoiceScore[outcome][cpu]
	}
	return playerScore
}

var shapeScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var roundScore = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 0,
		"C": 6,
	},
	"Y": {
		"A": 6,
		"B": 3,
		"C": 0,
	},
	"Z": {
		"A": 0,
		"B": 6,
		"C": 3,
	},
}

var outcomeScore = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

var outcomeToChoiceScore = map[string]map[string]int{
	"X": {
		"A": 3,
		"B": 1,
		"C": 2,
	},
	"Y": {
		"A": 1,
		"B": 2,
		"C": 3,
	},
	"Z": {
		"A": 2,
		"B": 3,
		"C": 1,
	},
}
