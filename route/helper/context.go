package routeHelper

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func UpdateRequestMetaDataInContext(ginContext *gin.Context) {
	requestMetaData := GetRequestMetadataFromContext(ginContext)
	requestMetaData.LatencyInNanoSecond = helper.GetUnixTimeInNanoSecond() - requestMetaData.StartEpoch
}
func GetHttpResponseFromContext(ginContext *gin.Context) model.HttpResponse {
	if dataInterface, ok := ginContext.Get(constant.ContextHttpResponse); ok {
		if metaData, ok := dataInterface.(model.HttpResponse); ok {
			return metaData
		}
	}
	//log this failure here
	return model.HttpResponse{}
}

func SetHttpReponseInContext(ginContext *gin.Context, response model.HttpResponse) {
	ginContext.Set(constant.ContextHttpResponse, response)
}

// GetRequestMetadata retrieves the request metadata from the Gin context.
func GetRequestMetadataFromContext(ginContext *gin.Context) *model.RequestMetaData {
	if data, exists := ginContext.Get(constant.ContextRequestMetaData); exists {
		if requestMeta, ok := data.(*model.RequestMetaData); ok {
			return requestMeta
		}
	}
	return BuildAndSetRequestMetaInContext(ginContext)
}

func BuildAndSetRequestMetaInContext(ctx *gin.Context) *model.RequestMetaData {
	request := ctx.Request

	requestMeta := &model.RequestMetaData{
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
	defer ctx.Set(constant.ContextRequestMetaData, requestMeta) //ensure that the context is always set even if an error occurs.
	helper.LogDetails(logrus.InfoLevel, constant.RequestReceivedMessage, *requestMeta)

	return requestMeta
}

func getRequestID() string {

	uuid := uuid.New()
	requestID := fmt.Sprintf("%v_%v_%v_%v", helper.GetApplicationConfiguration().ApplicationID, carbon.Now().DayOfYear(), carbon.Now().Hour(), uuid)

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
