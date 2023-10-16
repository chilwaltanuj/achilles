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
)

func RenderJsonResponse(ginContext *gin.Context) {
	response := GetHttpReponseFromContext(ginContext)
	ginContext.JSON(response.Status, response)
}

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
func GetRequestMetadataFromContext(c *gin.Context) model.RequestMetaData {
	data, exists := c.Get(constant.ContextRequestMetaData)
	if !exists {
		return BuildAndSetRequestMetaInContext(c)
	}
	if requestMeta, ok := data.(model.RequestMetaData); ok {
		return requestMeta
	}
	return BuildAndSetRequestMetaInContext(c)
}

func BuildAndSetRequestMetaInContext(ctx *gin.Context) model.RequestMetaData {
	request := ctx.Request

	requestMeta := model.RequestMetaData{
		ID:         getRequestID(),
		StatusCode: ctx.Writer.Status(),
		UserAgent:  request.UserAgent(),
		HttpMethod: request.Method,
		URL:        getCompleteURLFromRequest(request),
		Query:      request.URL.Query(),
		IP:         request.RemoteAddr,
		StartEpoch: helper.GetUnixTimeInNanoSecond(),
	}
	ctx.Set(constant.ContextRequestMetaData, requestMeta)

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
	serviceID := "001"
	uuid := uuid.New()
	// Generate a unique request ID (e.g., a UUID)
	requestID := fmt.Sprintf("%v_%v_%v_%v", serviceID, time.Now().YearDay(), time.Now().UTC().Hour(), uuid)

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
