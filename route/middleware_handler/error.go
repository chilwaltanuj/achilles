package middlewareHandler

import (
	"achilles/constant"
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func MethodNotSupported() gin.HandlerFunc {
	return func(gincontext *gin.Context) {
		routeHelper.SetSuccessResponseWithOnlyMessage(gincontext, constant.HttpMethodNotSUpported)
		RenderResponse(gincontext)
	}
}

func RouteNotSupported(gincontext *gin.Context) {
	routeHelper.SetSuccessResponseWithOnlyMessage(gincontext, constant.HttpRouteNotFound)
	RenderResponse(gincontext)
}
