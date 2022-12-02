package two

import (
	"bufio"
	"fmt"
	"os"
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
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resolveTaskOne(tt.args.input); got != tt.want {
				t.Errorf("resolveTaskOne() = %v, want %v", got, tt.want)
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
			if got := resolveTaskTwo(tt.args.input); got != tt.want {
				t.Errorf("resolveTaskTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSingleResult(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First example",
			args: args{input: []string{
				"A", "Y",
			}},
			want: 8,
		},
		{
			name: "Second example",
			args: args{input: []string{
				"B", "X",
			}},
			want: 1,
		},
		{
			name: "Third example",
			args: args{input: []string{
				"C", "Z",
			}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSingleResult(tt.args.input); got != tt.want {
				t.Errorf("getSingleResult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWinPoints(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First example",
			args: args{input: []string{
				"A", "Y",
			}},
			want: 6,
		},
		{
			name: "Second example",
			args: args{input: []string{
				"B", "X",
			}},
			want: 0,
		},
		{
			name: "Third example",
			args: args{input: []string{
				"C", "Z",
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWinPoints(tt.args.input); got != tt.want {
				t.Errorf("getWinPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
