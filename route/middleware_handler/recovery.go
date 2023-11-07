package middlewareHandler

import (
	"fmt"
	"net/http"

	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

// Recovery is a middleware that recovers from any panics.
func Recovery() gin.HandlerFunc {
	return RecoveryMiddleware()
}

// RecoveryMiddleware is the core recovery middleware function.
func RecoveryMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer handlePanic(ginContext)
		ginContext.Next()
	}
}

func handlePanic(ginContext *gin.Context) {
	if r := recover(); r != nil {
		responseData := model.HttpResponseData{
			Success: false,
			Status:  http.StatusOK,
			Message: constant.PanicRecovery,
		}
		routeHelper.BuildAndSetHttpResponseInContext(ginContext, responseData)
		routeHelper.UpdateRequestMetaDataInContext(ginContext)
		routeHelper.RenderJsonResponse(ginContext)

		errorMessage := fmt.Sprintf("%+v", r)
		helper.LogMessageWithStackTrace(errorMessage)
	}
}
