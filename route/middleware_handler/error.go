package middlewareHandler

import (
	"achilles/constant"
	routeHelper "achilles/route/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func MethodNotSupported() gin.HandlerFunc {
func MethodNotSupported(gincontext *gin.Context) {
	routeHelper.SetResponseWithMessageAndStatusCode(gincontext, constant.HttpMethodNotSUpported, http.StatusMethodNotAllowed)
	RenderResponse(gincontext)
}

//}

func RouteNotSupported(gincontext *gin.Context) {
	routeHelper.SetResponseWithMessageAndStatusCode(gincontext, constant.HttpRouteNotFound, http.StatusNotFound)
	RenderResponse(gincontext)
}
