package days

import "github.com/gin-gonic/gin"

type Day struct {
	Name               string
	AbsoluteFilepath   string
	Task1              func([]string) string
	Task2              func([]string) string
	RouteSetupFunction func(*gin.Engine, []string)
}
