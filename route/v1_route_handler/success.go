package v1RouteHandler

import (
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func Success(ginContext *gin.Context) {
	routeHelper.SetSuccessResponseWithOnlyMessage(ginContext, "Fairy God mother has Granted your wish!")
}
