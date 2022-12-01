package one

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func getTestArray() []string {
	readFile, err := os.Open("./test.txt")
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

func TestResolveTaskOne(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test input",
			args: args{
				input: getTestArray(),
			},
			want: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveTaskOne(tt.args.input); got != tt.want {
				t.Errorf("ResolveTaskOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestResolveTaskTwo(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test input",
			args: args{
				input: getTestArray(),
			},
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResolveTaskTwo(tt.args.input); got != tt.want {
				t.Errorf("ResolveTaskTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMostCalories(t *testing.T) {
	type args struct {
		calories []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test input",
			args: args{
				[]int{6000, 4000, 11000, 24000, 10000},
			},
			want: 24000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMostCalories(tt.args.calories); got != tt.want {
				t.Errorf("getMostCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSumOfTopThreeCalories(t *testing.T) {
	type args struct {
		calories []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test input",
			args: args{
				[]int{6000, 4000, 11000, 24000, 10000},
			},
			want: 45000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfTopThreeCalories(tt.args.calories); got != tt.want {
				t.Errorf("getSumOfTopThreeCalories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getListOfElves(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test input",
			args: args{
				input: getTestArray(),
			},
			want: []int{6000, 4000, 11000, 24000, 10000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getListOfElves(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getListOfElves() = %v, want %v", got, tt.want)
			}
		})
	}
}
