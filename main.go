package main

import (
	"adventofcode/days/one"
	"bufio"
	"fmt"
	"os"
)

type task struct {
	filepath string
	name     string
	task1    func([]string) int
	task2    func([]string) int
}

func main() {
	tasks := []task{
		{
			filepath: "./days/one/input.txt",
			name:     "one",
			task1:    one.ResolveTaskOne,
			task2:    one.ResolveTaskTwo,
		},
	}

	for _, task := range tasks {
		input := readFileToStringArray(task.filepath)
		fmt.Printf("Result for day %v task 1: %v\n", task.name, task.task1(input))
		fmt.Printf("Result for day %v task 2: %v\n", task.name, task.task2(input))
	}

}

func readFileToStringArray(filePath string) []string {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	stringArray := []string{}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		stringArray = append(stringArray, fileScanner.Text())
	}
	return stringArray
}
