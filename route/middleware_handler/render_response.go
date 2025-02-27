package middlewareHandler

import (
	"achilles/constant"
	"achilles/helper"
	"net/http"

	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

// RenderResponse renders the HTTP response in the desired format based on the "Accept" header.
func RenderResponse(gincontext *gin.Context) {
	acceptHeader := gincontext.Request.Header.Get(constant.HeaderAccept)
	response := routeHelper.GetHttpResponseFromContext(gincontext)

	// Use the appropriate response format handler
	switch acceptHeader {
	case constant.ContentTypeXML:
		gincontext.XML(response.Status, response)
	default:
		gincontext.JSON(response.Status, response)
	}
	if response.Status == http.StatusInternalServerError {
		logMetadata := routeHelper.GetErrorMetadataFromContext(gincontext)
		helper.LogDetails(constant.LogLevelError, constant.HttpServerErrorPanic, logMetadata, response)
	} else {
		helper.LogDetails(constant.LogLevelInfo, constant.RenderResponseMessage, response)
	}
}
