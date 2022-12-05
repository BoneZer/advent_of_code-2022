package two

import (
	"adventofcode/days"
	"strconv"
	"strings"
)

const winPoints = 6
const drawPoints = 3
const loosePoints = 0

var selectionPoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var matchConstellations = map[string]map[string]int{
	"A": {
		"Y": winPoints,
		"X": drawPoints,
		"Z": loosePoints,
	},
	"B": {
		"Z": winPoints,
		"Y": drawPoints,
		"X": loosePoints,
	},
	"C": {
		"X": winPoints,
		"Z": drawPoints,
		"Y": loosePoints,
	},
}

var suggestionStrategy = map[string]int{
	"X": loosePoints,
	"Y": drawPoints,
	"Z": winPoints,
}

func GetTask() days.Day {
	return days.Day{
		Name:             "Two",
		AbsoluteFilepath: "./days/two/input.txt",
		Task1:            resolveTaskOne,
		Task2:            resolveTaskTwo,
	}
}

func resolveTaskOne(input []string) string {
	return strconv.Itoa(getCompleteResult(input, false))
}

func resolveTaskTwo(input []string) string {
	return strconv.Itoa(getCompleteResult(input, true))
}

func getCompleteResult(input []string, swapForStrategy bool) int {
	sumResult := 0

	for _, inputLine := range input {
		match := strings.Split(inputLine, " ")
		if swapForStrategy {
			match[1] = getStrategySelection(match)
		}
		sumResult += getSingleResult(match)
	}

	return sumResult
}

func getSingleResult(input []string) int {
	ownSelection := input[1]
	result := 0
	result += selectionPoints[ownSelection]
	result += getWinPoints(input)

	return result
}

func getWinPoints(input []string) int {
	return matchConstellations[input[0]][input[1]]
}

func getStrategySelection(input []string) string {
	opponentSelection := input[0]
	strategy := suggestionStrategy[input[1]]

	x := matchConstellations[opponentSelection]

	for i := range x {
		if x[i] == strategy {
			return i
		}
	}

	return ""
}
