package one

import (
	"fmt"
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
	topThreeCalories := []int{0, 0, 0}

	for _, calories := range inputCalories {
		if calories > topThreeCalories[0] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = topThreeCalories[0]
			topThreeCalories[0] = calories
		} else if calories > topThreeCalories[1] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = calories

		} else if calories > topThreeCalories[2] {
			topThreeCalories[2] = calories
		}
	}

	for _, calories := range topThreeCalories {
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
