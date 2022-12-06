package three

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
		want string
	}{
		{
			name: "Test input",
			args: args{
				input: getTestArray(),
			},
			want: "157",
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
		want string
	}{
		{
			name: "Test input",
			args: args{
				input: getTestArray(),
			},
			want: "70",
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

func Test_splitBackpackInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "First Test input",
			args: args{
				input: "vJrwpWtwJgWrhcsFMMfFFhFp",
			},
			want: []string{
				"vJrwpWtwJgWr", "hcsFMMfFFhFp",
			},
		},

		{
			name: "Second Test input",
			args: args{
				input: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			},
			want: []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		}, {
			name: "Third Test input",
			args: args{
				input: "PmmdzqPrVvPwwTWBwg",
			},
			want: []string{"PmmdzqPrV", "vPwwTWBwg"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitBackpackInput(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitBackpackInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDoubleItem(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "First Test input",
			args: args{
				input: []string{
					"vJrwpWtwJgWr", "hcsFMMfFFhFp",
				},
			},
			want: 'p',
		},

		{
			name: "Second Test input",
			args: args{
				input: []string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			},
			want: 'L',
		}, {
			name: "Third Test input",
			args: args{
				input: []string{"PmmdzqPrV", "vPwwTWBwg"},
			},
			want: 'P',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDoubleItem(tt.args.input); got != tt.want {
				t.Errorf("findDoubleItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPriorityPoints(t *testing.T) {
	type args struct {
		input rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "First test input",
			args: args{input: 'p'},
			want: 16,
		},
		{
			name: "Second test input",
			args: args{input: 'L'},
			want: 38,
		},
		{
			name: "Third test input",
			args: args{input: 'P'},
			want: 42,
		},
		{
			name: "Fourth test input",
			args: args{input: 'v'},
			want: 22,
		},
		{
			name: "Fifth test input",
			args: args{input: 't'},
			want: 20,
		},
		{
			name: "Sixth test input",
			args: args{input: 's'},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPriorityPoints(tt.args.input); got != tt.want {
				t.Errorf("getPriorityPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGroupsOfElves(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test input",
			args: args{
				[]string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg", "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			want: [][]string{
				{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"}, {"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGroupsOfElves(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGroupsOfElves() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBadgeOfElveGroup(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "First Test input",
			args: args{
				input: []string{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
			},
			want: 'r',
		},
		{
			name: "Second Test input",
			args: args{
				input: []string{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
			},
			want: 'Z',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBadgeOfElveGroup(tt.args.input); got != tt.want {
				t.Errorf("getBadgeOfElveGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
