package one

import (
	"fmt"
	"sort"
	"strconv"
)

func ResolveTaskOne(input []string) int {
	return getMostCalories(getListOfElves(input))
}

func ResolveTaskTwo(input []string) int {
	return getSumOfTopThreeCalories(getListOfElves(input))
}

func getMostCalories(inputCalories []int) int {
	mostCalories := 0
	for _, calories := range inputCalories {
		if calories > mostCalories {
			mostCalories = calories
		}
	}
	return mostCalories
}

func getSumOfTopThreeCalories(inputCalories []int) int {
	mostCalories := 0

	sort.Ints(inputCalories)

	for _, calories := range inputCalories[len(inputCalories)-3:] {
		mostCalories += calories
	}

	return mostCalories
}

func getListOfElves(input []string) []int {

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
