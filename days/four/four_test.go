package four

import (
	"adventofcode/days"
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

func Test_resolveTaskOne(t *testing.T) {
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
			want: 2,
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

func Test_resolveTaskTwo(t *testing.T) {
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
			want: 4,
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

func Test_splitElves(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "First test input",
			args: args{
				input: "2-4,6-8",
			},
			want: []string{"2-4", "6-8"},
		},
		{
			name: "Second test input",
			args: args{
				input: "2-3,4-5",
			},
			want: []string{"2-3", "4-5"},
		},
		{
			name: "Third test input",
			args: args{
				input: "5-7,7-9",
			},
			want: []string{"5-7", "7-9"},
		},
		{
			name: "Fourth test input",
			args: args{
				input: "2-8,3-7",
			},
			want: []string{"2-8", "3-7"},
		},
		{
			name: "Fifth test input",
			args: args{
				input: "6-6,4-6",
			},
			want: []string{"6-6", "4-6"},
		},
		{
			name: "Sixth test input",
			args: args{
				input: "2-6,4-8",
			},
			want: []string{"2-6", "4-8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitElves(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitElves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTask(t *testing.T) {
	tests := []struct {
		name string
		want days.Day
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTask(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMinAndMaxSections(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "First test input",
			args: args{
				input: "2-4",
			},
			want: []int{2, 4},
		},
		{
			name: "Second test input",
			args: args{
				input: "6-8",
			},
			want: []int{6, 8},
		},
		{
			name: "Third test input",
			args: args{
				input: "6-6",
			},
			want: []int{6, 6},
		},
		{
			name: "Fourth test input",
			args: args{
				input: "7-9",
			},
			want: []int{7, 9},
		},
		{
			name: "Fifth test input",
			args: args{
				input: "2-8",
			},
			want: []int{2, 8},
		},
		{
			name: "Own test one",
			args: args{
				input: "1-1",
			},
			want: []int{1, 1},
		},
		{
			name: "Own test two",
			args: args{
				input: "99-102",
			},
			want: []int{99, 102},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinAndMaxSections(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMinAndMaxSections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sectionsCompleteOverlapping(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First test input",
			args: args{
				input: [][]int{
					{2, 4}, {6, 8},
				},
			},
			want: false,
		},
		{
			name: "Second test input",
			args: args{
				input: [][]int{
					{5, 7}, {7, 9},
				},
			},
			want: false,
		},
		{
			name: "Third test input",
			args: args{
				input: [][]int{
					{6, 6}, {4, 6},
				},
			},
			want: true,
		},
		{
			name: "Fourth test input",
			args: args{
				input: [][]int{
					{2, 8}, {3, 7},
				},
			},
			want: true,
		},
		{
			name: "Edge case one",
			args: args{
				input: [][]int{
					{1, 4}, {1, 6},
				},
			},
			want: true,
		},
		{
			name: "Edge case two",
			args: args{
				input: [][]int{
					{2, 39}, {1, 39},
				},
			},
			want: true,
		},
		{
			name: "Edge case three",
			args: args{
				input: [][]int{
					{1, 2}, {3, 99},
				},
			},
			want: false,
		},
		{
			name: "Edge case four",
			args: args{
				input: [][]int{
					{1, 12}, {1, 4},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sectionsCompleteOverlapping(tt.args.input); got != tt.want {
				t.Errorf("sectionsCompleteOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sectionsOverlapping(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "First test input",
			args: args{
				input: [][]int{
					{2, 4}, {6, 8},
				},
			},
			want: false,
		},
		{
			name: "Second test input",
			args: args{
				input: [][]int{
					{2, 3}, {4, 5},
				},
			},
			want: false,
		},
		{
			name: "Third test input",
			args: args{
				input: [][]int{
					{5, 7}, {7, 9},
				},
			},
			want: true,
		},
		{
			name: "Fourth test input",
			args: args{
				input: [][]int{
					{2, 8}, {3, 7},
				},
			},
			want: true,
		},
		{
			name: "Fifth test input",
			args: args{
				input: [][]int{
					{6, 6}, {4, 6},
				},
			},
			want: true,
		},
		{
			name: "Sixth test input",
			args: args{
				input: [][]int{
					{2, 6}, {4, 8},
				},
			},
			want: true,
		},
		{
			name: "Edge case one",
			args: args{
				input: [][]int{
					{1, 4}, {1, 6},
				},
			},
			want: true,
		},
		{
			name: "Edge case two",
			args: args{
				input: [][]int{
					{2, 39}, {1, 39},
				},
			},
			want: true,
		},
		{
			name: "Edge case three",
			args: args{
				input: [][]int{
					{1, 2}, {3, 99},
				},
			},
			want: false,
		},
		{
			name: "Edge case four",
			args: args{
				input: [][]int{
					{1, 12}, {1, 4},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sectionsOverlapping(tt.args.input); got != tt.want {
				t.Errorf("sectionsOverlapping() = %v, want %v", got, tt.want)
			}
		})
	}
}
