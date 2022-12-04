package three

import (
	"adventofcode/days"

	"github.com/gin-gonic/gin"
)

func GetTask() days.Day {
	return days.Day{
		Name:               "Three",
		AbsoluteFilepath:   "./days/three/input.txt",
		Task1:              resolveTaskOne,
		Task2:              resolveTaskTwo,
		RouteSetupFunction: setUpRoutes,
	}
}

func setUpRoutes(router *gin.Engine, input []string) {
}

func resolveTaskOne(input []string) int {
	itemSum := 0

	for _, backpack := range input {
		comparsements := splitBackpackInput(backpack)
		doubleItem := findDoubleItem(comparsements)
		itemSum += getPriorityPoints(doubleItem)
	}

	return itemSum
}

func resolveTaskTwo(input []string) int {
	itemSum := 0
	groupsOfElves := getGroupsOfElves(input)
	for _, groupOfElves := range groupsOfElves {
		badge := getBadgeOfElveGroup(groupOfElves)
		itemSum += getPriorityPoints(badge)
	}

	return itemSum
}

func splitBackpackInput(input string) []string {
	return []string{input[0 : len(input)/2], input[len(input)/2:]}
}

func findDoubleItem(input []string) rune {
	for _, item := range input[0] {
		for _, secondItem := range input[1] {
			if item == secondItem {
				return item
			}
		}
	}
	return 0
}

func getPriorityPoints(input rune) int {
	if input > 96 && input < 123 {
		return int(input) - 96
	} else if input > 64 && input < 91 {
		return int(input) - 38
	}
	return 0
}

func getGroupsOfElves(input []string) [][]string {
	groupOfElves := [][]string{}

	for i := 0; i < len(input)-2; i += 3 {
		group := []string{input[i], input[i+1], input[i+2]}
		groupOfElves = append(groupOfElves, group)
	}

	return groupOfElves
}

func getBadgeOfElveGroup(input []string) rune {
	for _, firstElf := range input[0] {
		for _, secondElf := range input[1] {
			for _, thirdElf := range input[2] {
				if firstElf == secondElf && secondElf == thirdElf {
					return firstElf
				}
			}
		}
	}

	return 0
}
