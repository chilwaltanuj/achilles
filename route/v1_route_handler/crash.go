package v1RouteHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Crash(context *gin.Context) {
	dependencyContainer.LogWriter.Info(" I am about to crash! Yipeee")

	divisor := 0
	dividend := 1
	dependencyContainer.LogWriter.Info("is this happening")
	fmt.Println(dividend / divisor)
}
