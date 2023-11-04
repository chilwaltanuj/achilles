package routeHelper

import (
	"achilles/constant"
	"achilles/model"

	"github.com/gin-gonic/gin"
)

func SetSuccessResponse(ginContext *gin.Context, data model.ResponseData) {
	responseData := model.HttpResponseData{
		Success:      true,
		Status:       200,
		Message:      constant.HttpSuccess,
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
