package days

import "github.com/gin-gonic/gin"

type Day struct {
	Name               string
	AbsoluteFilepath   string
	Task1              func([]string) int
	Task2              func([]string) int
	RouteSetupFunction func(*gin.Engine, []string)
}
