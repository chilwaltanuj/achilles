package middlewareHandler

import (
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func MethodNotSupported() gin.HandlerFunc {
	return func(c *gin.Context) {
		routeHelper.SetSuccessResponseWithOnlyMessage(c, "check the method used. Not supported")
		routeHelper.RenderJsonResponse(c)
	}
}

func RouteNotSupported(c *gin.Context) {
	routeHelper.SetSuccessResponseWithOnlyMessage(c, "check the resource path")
	routeHelper.RenderJsonResponse(c)
}
