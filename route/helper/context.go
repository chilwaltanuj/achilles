package routeHelper

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateRequestMetaDataInContext(ginContext *gin.Context) {
	requestMetaData := GetRequestMetadataFromContext(ginContext)
	requestMetaData.LatencyInNanoSecond = helper.GetUnixTimeInNanoSecond() - requestMetaData.StartEpoch
	ginContext.Set(constant.ContextRequestMetaData, requestMetaData)
}
func GetHttpReponseFromContext(ginContext *gin.Context) model.HttpResponse {
	if dataInterface, ok := ginContext.Get(constant.ContextHttpResponse); ok {
		if metaData, ok := dataInterface.(model.HttpResponse); ok {
			return metaData
		}
	}
	//log this failure here
	return model.HttpResponse{}
}

func GetHttpReponseDataFromContext(ginContext *gin.Context) model.HttpResponseData {
	if dataInterface, ok := ginContext.Get(constant.ContextHttpResponseData); ok {
		if metaData, ok := dataInterface.(model.HttpResponseData); ok {
			return metaData
		}
	}
	//log this failure here
	return model.HttpResponseData{}
}

func SetHttpReponseInContext(ginContext *gin.Context, response model.HttpResponse) {
	ginContext.Set(constant.ContextHttpResponse, response)
}

// GetRequestMetadata retrieves the request metadata from the Gin context.
func GetRequestMetadataFromContext(ginContext *gin.Context) model.RequestMetaData {
	data, exists := ginContext.Get(constant.ContextRequestMetaData)
	if !exists { //request received for the first time. build it first and then return it
		return BuildAndSetRequestMetaInContext(ginContext)
	}
	if requestMeta, ok := data.(model.RequestMetaData); ok {
		return requestMeta
	}
	return BuildAndSetRequestMetaInContext(ginContext)
}

func BuildAndSetRequestMetaInContext(ctx *gin.Context) model.RequestMetaData {
	request := ctx.Request

	requestMeta := model.RequestMetaData{
		URL:           getCompleteURLFromRequest(request),
		HttpMethod:    request.Method,
		StatusCode:    ctx.Writer.Status(),
		Query:         request.URL.Query(),
		ID:            getRequestID(),
		IP:            request.RemoteAddr,
		StartEpoch:    helper.GetUnixTimeInNanoSecond(),
		UserAgent:     request.UserAgent(),
		ApplicationID: helper.GetApplicationConfiguration().ApplicationID,
		Application:   helper.GetApplicationConfiguration().Application,
	}
	ctx.Set(constant.ContextRequestMetaData, requestMeta)
	helper.LogDetails(logrus.InfoLevel, constant.RequestReceivedMessage, requestMeta)
	return requestMeta
}

func FetchRequestMetaDataFromContext(ginContext *gin.Context) model.RequestMetaData {
	return FetchHttpReponseFromContext(ginContext).MetaData
}

func FetchHttpReponseFromContext(ginContext *gin.Context) model.HttpResponse {
	if dataInterface, ok := ginContext.Get(constant.ContextHttpResponse); ok {
		if metaData, ok := dataInterface.(model.HttpResponse); ok {
			return metaData
		}
	}
	//log this failure here
	return model.HttpResponse{}
}

func getRequestID() string {

	uuid := uuid.New()
	requestID := fmt.Sprintf("%v_%v_%v_%v", helper.GetApplicationConfiguration().ApplicationID, time.Now().YearDay(), time.Now().UTC().Hour(), uuid)

	return requestID
}

func BuildAndSetHttpResponseInContext(ginContext *gin.Context, responseData model.HttpResponseData) {
	httpResponse := model.HttpResponse{
		HttpResponseData: responseData,
		MetaData:         GetRequestMetadataFromContext(ginContext),
	}

	SetHttpReponseInContext(ginContext, httpResponse)
}

func getCompleteURLFromRequest(httpRequest *http.Request) string {
	scheme := "http"
	if httpRequest.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + httpRequest.Host + httpRequest.RequestURI
}
