package v1RouteHandler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Crash(context *gin.Context) {
	divisor := 0
	dividend := 1
	fmt.Println(dividend / divisor)
}
