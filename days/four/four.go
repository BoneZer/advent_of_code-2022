package four

import (
	"adventofcode/days"
	"strconv"
	"strings"
)

func GetTask() days.Day {
	return days.Day{
		Name:             "Two",
		AbsoluteFilepath: "./days/four/input.txt",
		Task1:            resolveTaskOne,
		Task2:            resolveTaskTwo,
	}
}

func resolveTaskOne(input []string) int {
	overlappingsSum := 0

	for _, line := range input {
		elves := splitElves(line)
		sections := [][]int{}
		for _, elf := range elves {
			sections = append(sections, getMinAndMaxSections(elf))
		}
		if sectionsCompleteOverlapping(sections) {
			overlappingsSum++
		}
	}

	return overlappingsSum
}

func resolveTaskTwo(input []string) int {
	overlappingsSum := 0

	for _, line := range input {
		elves := splitElves(line)
		sections := [][]int{}
		for _, elf := range elves {
			sections = append(sections, getMinAndMaxSections(elf))
		}
		if sectionsOverlapping(sections) {
			overlappingsSum++
		}
	}

	return overlappingsSum
}

func splitElves(input string) []string {
	return strings.Split(input, ",")
}

func getMinAndMaxSections(input string) []int {
	sectionRange := strings.Split(input, "-")
	minSection, _ := strconv.Atoi(sectionRange[0])
	maxSection, _ := strconv.Atoi(sectionRange[1])
	sections := []int{minSection, maxSection}

	return sections
}

func sectionsCompleteOverlapping(input [][]int) bool {
	minFirstSection := input[0][0]
	maxFirstSection := input[0][1]
	minSecondSection := input[1][0]
	maxSecondSection := input[1][1]

	if minFirstSection < minSecondSection {
		if maxFirstSection > maxSecondSection || maxFirstSection == maxSecondSection {
			return true
		}
	} else if minFirstSection == minSecondSection {
		return true
	} else {
		if maxSecondSection > maxFirstSection || maxSecondSection == maxFirstSection {
			return true
		}
	}

	return false
}

func sectionsOverlapping(input [][]int) bool {
	minFirstSection := input[0][0]
	maxFirstSection := input[0][1]
	minSecondSection := input[1][0]
	maxSecondSection := input[1][1]

	if minFirstSection < minSecondSection {
		if maxFirstSection >= minSecondSection {
			return true
		}
	} else if minSecondSection < minFirstSection {
		if maxSecondSection >= minFirstSection {
			return true
		}
	} else if minFirstSection == minSecondSection {
		return true
	}

	return false
}
