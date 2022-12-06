package five

import (
	"adventofcode/days"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetTask() days.Day {
	return days.Day{
		Name:               "Five",
		AbsoluteFilepath:   "./days/five/input.txt",
		Task1:              resolveTaskOne,
		Task2:              resolveTaskTwo,
		RouteSetupFunction: setUpRoutes,
	}
}

func setUpRoutes(router *gin.Engine, input []string) {
}

func resolveTaskOne(input []string) string {
	startContainer, commands := splitDataInput(input)

	containerState := getStartArray(startContainer)
	for _, command := range commands {
		containerState = removeEmptyColumns(containerState)
		containerState = moveContainers(containerState, command)
	}

	result := ""
	for _, containerStack := range containerState {
		result += containerStack[0]
	}

	return result
}

func resolveTaskTwo(input []string) string {
	return "2"
}

func splitDataInput(input []string) ([]string, []string) {
	splitIndex := 0

	for i := range input {
		if input[i] == "" {
			splitIndex = i
		}
	}

	return input[0:splitIndex], input[splitIndex+1:]
}

func getStartArray(input []string) [][]string {
	startContainerLines := [][]string{}

	for i := 0; i < len(input)-1; i++ {
		startContainerLines = append(startContainerLines, getContainerLine(input[i]))
	}

	startContainerStacks := removeEmptyColumns(getEmptyStacks(len(startContainerLines[0]), len(startContainerLines)))

	for i := 0; i < len(startContainerLines[0]); i++ {
		for j := 0; j < len(startContainerLines); j++ {
			startContainerStacks[i] = append(startContainerStacks[i], startContainerLines[j][i])
		}
	}

	return removeEmptyColumns(startContainerStacks)
}

func getContainerLine(textLine string) []string {
	containerLine := []string{}
	splittedInput := strings.Split(textLine, "")

	for i := 1; i < len(splittedInput); i += 4 {
		if splittedInput[i] == " " {
			containerLine = append(containerLine, "")
		} else {
			containerLine = append(containerLine, splittedInput[i])
		}
	}

	return containerLine
}

func getEmptyStacks(width int, height int) [][]string {
	emptyStack := [][]string{}
	for i := 0; i < width; i++ {
		emptyStack = append(emptyStack, make([]string, height))
	}

	return emptyStack
}

func moveContainers(containerState [][]string, moveCommand string) [][]string {
	re := regexp.MustCompile("[0-9]+")
	commandNumbers := re.FindAllString(moveCommand, -1)
	moveCount, _ := strconv.Atoi(commandNumbers[0])
	from, _ := strconv.Atoi(commandNumbers[1])
	to, _ := strconv.Atoi(commandNumbers[2])

	for i := 0; i < moveCount; i++ {
		containerToMove := containerState[from-1][0]
		containerState[from-1] = containerState[from-1][1:]
		containerState[to-1] = append([]string{containerToMove}, containerState[to-1]...)
	}

	return containerState
}

func removeEmptyColumns(containerState [][]string) [][]string {
	for i := 0; i < len(containerState); i++ {
		indexWithLetter := len(containerState[i])
		for j := 0; j < len(containerState[i]); j++ {
			if containerState[i][j] != "" {
				indexWithLetter = j
				break
			}
		}
		containerState[i] = containerState[i][indexWithLetter:len(containerState[i])]
	}

	return containerState
}
