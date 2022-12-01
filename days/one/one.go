package one

import (
	"fmt"
	"strconv"
)

type elf struct {
	calories int
}

func ResolveTaskOne(input []string) int {
	return getMostCalories(getListOfElves(input))
}

func ResolveTaskTwo(input []string) int {
	return getSumOfTopThreeCalories(getListOfElves(input))
}

func getMostCalories(elves []elf) int {
	mostCalories := 0
	for _, elf := range elves {
		if elf.calories > mostCalories {
			mostCalories = elf.calories
		}
	}
	return mostCalories
}

func getSumOfTopThreeCalories(elves []elf) int {
	mostCalories := 0
	topThreeCalories := []int{0, 0, 0}

	for _, elf := range elves {
		if elf.calories > topThreeCalories[0] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = topThreeCalories[0]
			topThreeCalories[0] = elf.calories
		} else if elf.calories > topThreeCalories[1] {
			topThreeCalories[2] = topThreeCalories[1]
			topThreeCalories[1] = elf.calories

		} else if elf.calories > topThreeCalories[2] {
			topThreeCalories[2] = elf.calories
		}
	}

	for _, calories := range topThreeCalories {
		mostCalories += calories
	}

	return mostCalories
}

func getListOfElves(input []string) []elf {
	elves := []elf{}

	elf := elf{calories: 0}

	for _, line := range input {
		if line == "" {
			elves = append(elves, elf)
			elf.calories = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			elf.calories += calories
		}
	}
	elves = append(elves, elf)

	return elves
}
