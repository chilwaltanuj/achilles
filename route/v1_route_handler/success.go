package v1RouteHandler

import (
	"achilles/constant"
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func Success(ginContext *gin.Context) {
	routeHelper.SetSuccessResponseWithOnlyMessage(ginContext, constant.HttpSuccess)
}
