package five

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
			want: "CMZ",
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
			want: "b",
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

func Test_getStartArray(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "First test input",
			args: args{
				input: []string{
					"    [D]    ",
					"[N] [C]    ",
					"[Z] [M] [P]",
					" 1   2   3 ",
				},
			},
			want: [][]string{
				{"N", "Z"}, {"D", "C", "M"}, {"P"},
			},
		},
		{
			name: "First fail test",
			args: args{
				input: []string{"            [C]         [N] [R]    ", "[J] [T]     [H]         [P] [L]    ", "[F] [S] [T] [B]         [M] [D]    ", "[C] [L] [J] [Z] [S]     [L] [B]    ", "[N] [Q] [G] [J] [J]     [F] [F] [R]", "[D] [V] [B] [L] [B] [Q] [D] [M] [T]", "[B] [Z] [Z] [T] [V] [S] [V] [S] [D]", "[W] [P] [P] [D] [G] [P] [B] [P] [V]", " 1   2   3   4   5   6   7   8   9 "},
			},
			want: [][]string{
				{"J", "F", "C", "N", "D", "B", "W"}, {"T", "S", "L", "Q", "V", "Z", "P"},
				{"T", "J", "G", "B", "Z", "P"}, {"C", "H", "B", "Z", "J", "L", "T", "D"},
				{"S", "J", "B", "V", "G"}, {"Q", "S", "P"}, {"N", "P", "M", "L", "F", "D", "V", "B"},
				{"R", "L", "D", "B", "F", "M", "S", "P"}, {"R", "T", "D", "V"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStartArray(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getStartArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveContainers(t *testing.T) {
	type args struct {
		containerState [][]string
		moveCommand    string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Second move",
			args: args{
				containerState: [][]string{
					{"N", "Z"}, {"D", "C", "M"}, {"P"},
				},
				moveCommand: "move 1 from 2 to 1",
			},
			want: [][]string{
				{"D", "N", "Z"}, {"C", "M"}, {"P"},
			},
		},
		{
			name: "First move",
			args: args{
				containerState: [][]string{
					{"D", "N", "Z"}, {"C", "M"}, {"P"},
				},
				moveCommand: "move 3 from 1 to 3",
			},
			want: [][]string{
				{}, {"C", "M"}, {"Z", "N", "D", "P"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveContainers(tt.args.containerState, tt.args.moveCommand); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveContainers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEmptyStacks(t *testing.T) {
	type args struct {
		width  int
		height int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "First test",
			args: args{
				width:  1,
				height: 1,
			},
			want: [][]string{{""}},
		},
		{
			name: "Second test",
			args: args{
				width:  2,
				height: 2,
			},
			want: [][]string{{"", ""}, {"", ""}},
		},
		{
			name: "Third test",
			args: args{
				width:  3,
				height: 2,
			},
			want: [][]string{{"", ""}, {"", ""}, {"", ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEmptyStacks(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEmptyStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeEmptyColumns(t *testing.T) {
	type args struct {
		containerState [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "First test",
			args: args{
				containerState: [][]string{{"", "N", "Z"}, {"D", "C", "M"}, {"", "", "P"}},
			},
			want: [][]string{{"N", "Z"}, {"D", "C", "M"}, {"P"}},
		},
		{
			name: "Second test",
			args: args{
				containerState: [][]string{{"", "", "", ""}, {"", "", "C", "M"}, {"Z", "M", "P", "D"}},
			},
			want: [][]string{{}, {"C", "M"}, {"Z", "M", "P", "D"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeEmptyColumns(tt.args.containerState); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeEmptyColumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitDataInput(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 []string
	}{
		{
			name: "First test input",
			args: args{
				input: []string{
					"    [D]    ",
					"[N] [C]    ",
					"[Z] [M] [P]",
					" 1   2   3 ",
					"",
					"move 1 from 2 to 1",
					"move 3 from 1 to 3",
				},
			},
			want: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3 ",
			},
			want1: []string{
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := splitDataInput(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitDataInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("splitDataInput() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getContainerLine(t *testing.T) {
	type args struct {
		textLine string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Fail Test 1",
			args: args{
				textLine: "    [D]    ",
			},
			want: []string{"", "D", ""},
		},
		{
			name: "Fail Test 2",
			args: args{
				textLine: "[N] [C]    ",
			},
			want: []string{"N", "C", ""},
		},
		{
			name: "Fail Test 3",
			args: args{
				textLine: "[Z] [M] [P]",
			},
			want: []string{"Z", "M", "P"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContainerLine(tt.args.textLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getContainerLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
