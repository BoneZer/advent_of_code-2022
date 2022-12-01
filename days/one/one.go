package one

import (
	"adventofcode/days"
	"fmt"
	"sort"
	"strconv"
)

func GetTask() days.Day {
	return days.Day{
		Name:             "One",
		AbsoluteFilepath: "./days/one/input.txt",
		Task1:            resolveTaskOne,
		Task2:            resolveTaskTwo,
	}
}

func resolveTaskOne(input []string) int {
	return getMostCalories(getListOfCalories(input))
}

func resolveTaskTwo(input []string) int {
	return getSumOfTopThreeCalories(getListOfCalories(input))
}

func getMostCalories(inputCalories []int) int {
	sort.Ints(inputCalories)
	return inputCalories[len(inputCalories)-1]
}

func getSumOfTopThreeCalories(inputCalories []int) int {
	mostCalories := 0

	sort.Ints(inputCalories)

	for _, calories := range inputCalories[len(inputCalories)-3:] {
		mostCalories += calories
	}

	return mostCalories
}

func getListOfCalories(input []string) []int {
	caloriesArray := []int{}
	caloriesSum := 0

	for _, line := range input {
		if line == "" {
			caloriesArray = append(caloriesArray, caloriesSum)
			caloriesSum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			caloriesSum += calories
		}
	}
	caloriesArray = append(caloriesArray, caloriesSum)

	return caloriesArray
}
