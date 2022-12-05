package main

import (
	day "adventofcode/days"
	"adventofcode/days/five"
	"adventofcode/days/four"
	"adventofcode/days/one"
	"adventofcode/days/three"
	"adventofcode/days/two"
	"bufio"
	"fmt"
	"os"
)

func main() {

	days := []day.Day{
		one.GetTask(),
		two.GetTask(),
		three.GetTask(),
		four.GetTask(),
		five.GetTask(),
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
