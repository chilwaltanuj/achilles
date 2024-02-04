package routeHelper

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetSuccessResponse(ginContext *gin.Context, data *model.ResponseData) {
	setResponse(ginContext, http.StatusOK, constant.HttpOk, data, true)
}

func SetSuccessResponseWithOnlyMessage(ginContext *gin.Context, message string) {
	setResponse(ginContext, http.StatusOK, message, nil, true)
}

func SetFailureResponseWithStatusCode(ginContext *gin.Context, httpStatus int, err error) {
	message := getFailureMessageByStatusCode(httpStatus)
	setResponse(ginContext, httpStatus, message, nil, false)
	helper.LogDetails(constant.LogLevelError, message, err)
}

func SetResponseWithMessageAndStatusCode(ginContext *gin.Context, message string, httpStatus int) {
	setResponse(ginContext, httpStatus, message, nil, httpStatus < 400)
}

// Helper function to avoid code duplication and ensure the 'Success' flag aligns with the response status.
func setResponse(ginContext *gin.Context, status int, message string, data *model.ResponseData, success bool) {
	responseData := model.HttpResponseData{
		Success:      success,
		Status:       status,
		Message:      message,
		ResponseData: data,
	}
	BuildAndSetHttpResponseInContext(ginContext, responseData)
	ginContext.Status(status) // Ensure HTTP status code is set on the response as well.
}

// Extract the switch statement into its own function for cleaner code and easier maintenance.
func getFailureMessageByStatusCode(status int) string {
	switch status {
	case http.StatusFailedDependency:
		return constant.DependencyFailed
	case http.StatusMethodNotAllowed:
		return constant.HttpMethodNotSupported // Fixed typo in constant name
	case http.StatusNotFound:
		return constant.HttpRouteNotFound
	default:
		return "Something failed"
	}
}
