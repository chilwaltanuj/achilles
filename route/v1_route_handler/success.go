package v1RouteHandler

import (
	"achilles/helper"
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func Success(ginContext *gin.Context) {
	helper.GetGlobalLogger().Info("Request Received")

	routeHelper.SetSuccessResponseWithOnlyMessage(ginContext, "Fairy God mother has Granted your wish!")
}
