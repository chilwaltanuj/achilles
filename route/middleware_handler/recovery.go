package middlewareHandler

import (
	"fmt"
	"net/http"
	"strings"

	"achilles/constant"
	"achilles/model"
	routeHelper "achilles/route/helper"

	"runtime/debug"

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
		stack := string(debug.Stack())
		stackLines := strings.Split(stack, "\n")
		// Skip the first 5 lines of the stack trace
		trimmedStack := strings.Join(stackLines[5:], "\n")

		responseData := model.HttpResponseData{
			Success: false,
			Status:  http.StatusInternalServerError,
			Message: constant.HttpServerErrorPanic,
		}
		routeHelper.BuildAndSetHttpResponseInContext(ginContext, responseData)

		// Create structured error metadata and store it in the context
		errorMetadata := &model.ErrorMetadata{
			Error:      fmt.Sprintf("%+v", r),
			Stacktrace: trimmedStack,
		}
		routeHelper.SeErrorMetadataInContext(ginContext, *errorMetadata)

		RenderResponse(ginContext)
	}
}
