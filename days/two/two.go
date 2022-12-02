package two

import (
	"adventofcode/days"
	"strings"
)

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
	return 2
}

func getSingleResult(input []string) int {
	ownSelection := input[1]
	result := 0
	if ownSelection == "X" {
		result = 1
	} else if ownSelection == "Y" {
		result = 2
	} else if ownSelection == "Z" {
		result = 3
	}
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
