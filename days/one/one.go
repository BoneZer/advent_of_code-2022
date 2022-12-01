package one

import (
	"fmt"
	"sort"
	"strconv"
)

func ResolveTaskOne(input []string) int {
	return getMostCalories(getListOfCalories(input))
}

func ResolveTaskTwo(input []string) int {
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
