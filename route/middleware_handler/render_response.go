package middlewareHandler

import (
	"achilles/constant"
	"achilles/helper"

	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// RenderResponse renders the HTTP response in the desired format based on the "Accept" header.
func RenderResponse(gincontext *gin.Context) {
	acceptHeader := gincontext.Request.Header.Get(constant.HeaderAccept)
	response := routeHelper.GetHttpReponseFromContext(gincontext)

	// Use the appropriate response format handler
	switch acceptHeader {
	case constant.ContentTypeXML:
		gincontext.XML(response.Status, response)
	default:
		gincontext.JSON(response.Status, response)
	}

	helper.LogDetails(logrus.InfoLevel, "Response rendered", response)
}
