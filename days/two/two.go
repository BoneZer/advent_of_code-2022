package two

import (
	"adventofcode/days"
	"strings"
)

var selectionPoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func GetTask() days.Day {
	return days.Day{
		Name:             "Two",
		AbsoluteFilepath: "./days/two/input.txt",
		Task1:            resolveTaskOne,
		Task2:            resolveTaskTwo,
	}
}

func resolveTaskOne(input []string) int {
	sumResult := 0

	for _, inputLine := range input {
		match := strings.Split(inputLine, " ")
		sumResult += getSingleResult(match)
	}

	return sumResult
}

func resolveTaskTwo(input []string) int {
	sumResult := 0

	for _, inputLine := range input {
		match := strings.Split(inputLine, " ")
		match[1] = getStrategySelection(match)
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
	opponentSelection := input[0]
	ownSelection := input[1]

	if opponentSelection == "A" {
		if ownSelection == "X" {
			return 3
		}
		if ownSelection == "Y" {
			return 6
		}
		if ownSelection == "Z" {
			return 0
		}
	}
	if opponentSelection == "B" {
		if ownSelection == "X" {
			return 0
		}
		if ownSelection == "Y" {
			return 3
		}
		if ownSelection == "Z" {
			return 6
		}
	}
	if opponentSelection == "C" {
		if ownSelection == "X" {
			return 6
		}
		if ownSelection == "Y" {
			return 0
		}
		if ownSelection == "Z" {
			return 3
		}
	}

	return 0
}

func getStrategySelection(input []string) string {
	opponentSelection := input[0]
	ownSelection := input[1]

	if opponentSelection == "A" {
		if ownSelection == "X" {
			return "Z"
		}
		if ownSelection == "Y" {
			return "X"
		}
		if ownSelection == "Z" {
			return "Y"
		}
	}
	if opponentSelection == "B" {
		if ownSelection == "X" {
			return "X"
		}
		if ownSelection == "Y" {
			return "Y"
		}
		if ownSelection == "Z" {
			return "Z"
		}
	}
	if opponentSelection == "C" {
		if ownSelection == "X" {
			return "Y"
		}
		if ownSelection == "Y" {
			return "Z"
		}
		if ownSelection == "Z" {
			return "X"
		}
	}

	return ""
}
