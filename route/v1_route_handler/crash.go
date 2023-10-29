package v1RouteHandler

import (
	"achilles/helper"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Crash(context *gin.Context) {
	helper.GetGlobalLogger().Info(" I am about to crash! Yipeee")

	divisor := 0
	dividend := 1
	helper.GetGlobalLogger().Info("is this happening")
	fmt.Println(dividend / divisor)
}
