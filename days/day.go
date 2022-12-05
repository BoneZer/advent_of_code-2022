package days

type Day struct {
	Name             string
	AbsoluteFilepath string
	Task1            func([]string) string
	Task2            func([]string) string
}
