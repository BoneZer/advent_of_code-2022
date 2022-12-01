package days

type Day struct {
	Name             string
	AbsoluteFilepath string
	Task1            func([]string) int
	Task2            func([]string) int
}
