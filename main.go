package main

import (
	day "adventofcode/days"
	"adventofcode/days/four"
	"adventofcode/days/one"
	"adventofcode/days/three"
	"adventofcode/days/two"
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	days := []day.Day{
		one.GetTask(),
		two.GetTask(),
		three.GetTask(),
		four.GetTask(),
	}

	for _, day := range days {
		input := readFileToStringArray(day.AbsoluteFilepath)
		fmt.Printf("Result for day %v task 1: %v\n", day.Name, day.Task1(input))
		fmt.Printf("Result for day %v task 2: %v\n", day.Name, day.Task2(input))
		router.GET(strings.ToLower(day.Name)+"/input", getRoutesForDays(day))
		day.RouteSetupFunction(router, input)
	}

	router.Run("localhost:8080")

}

func getRoutesForDays(day day.Day) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, readFileToStringArray(day.AbsoluteFilepath))
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
