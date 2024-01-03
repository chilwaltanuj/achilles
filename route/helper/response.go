package routeHelper

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetSuccessResponse(ginContext *gin.Context, data *model.ResponseData) {
	responseData := model.HttpResponseData{
		Success:      true,
		Status:       200,
		Message:      constant.HttpOk,
		ResponseData: data,
	}
	BuildAndSetHttpResponseInContext(ginContext, responseData)
}

func SetSuccessResponseWithOnlyMessage(ginContext *gin.Context, message string) {
	responseData := model.HttpResponseData{
		Success: true,
		Status:  200,
		Message: message,
	}
	BuildAndSetHttpResponseInContext(ginContext, responseData)
}

func SetFailureResponseWithStatusCode(ginContext *gin.Context, httpStatus int, err error) {
	message := "something Failed"

	switch httpStatus {
	case http.StatusFailedDependency:
		message = constant.DependencyFailed
	case http.StatusMethodNotAllowed:
		message = constant.HttpMethodNotSUpported
	case http.StatusNotFound:
		message = constant.HttpRouteNotFound
	}
	responseData := model.HttpResponseData{
		Success: true,
		Status:  httpStatus,
		Message: message,
	}
	BuildAndSetHttpResponseInContext(ginContext, responseData)
	helper.LogDetails(logrus.ErrorLevel, message, err)
}

func SetResponseWithMessageAndStatusCode(ginContext *gin.Context, message string, httpStatus int) {
	responseData := model.HttpResponseData{
		Success: true,
		Status:  httpStatus,
		Message: message,
	}
	BuildAndSetHttpResponseInContext(ginContext, responseData)
}
