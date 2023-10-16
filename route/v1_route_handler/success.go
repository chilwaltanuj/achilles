package v1RouteHandler

import (
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Success(ginContext *gin.Context) {
	logrus.Infof("wow! invoked")

	routeHelper.SetSuccessResponseWithOnlyMessage(ginContext, "What is happening")
}
