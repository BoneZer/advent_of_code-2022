package main

import (
	day "adventofcode/days"
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

	days := []day.Day{
		one.GetTask(),
	}

	for _, day := range days {
		input := readFileToStringArray(day.AbsoluteFilepath)
		fmt.Printf("Result for day %v task 1: %v\n", day.Name, day.Task1(input))
		fmt.Printf("Result for day %v task 2: %v\n", day.Name, day.Task2(input))
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
