package v1RouteHandler

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"context"
	"net/http"
	"time"

	routeHelper "achilles/route/helper"

	"github.com/gin-gonic/gin"
)

func Joke(ginContext *gin.Context) {
	request := model.ClientHTTPRequest{
		Method:  http.MethodGet,
		URL:     constant.URL_Joke_Random,
		Headers: constant.Headers,
		Body:    nil,
	}
	//TODO - context timeout need to be configurable
	contextWithTimeout, cancel := context.WithTimeout(ginContext, 2000*time.Millisecond)
	defer cancel()

	if response := helper.Execute[model.Joke_Random](contextWithTimeout, request); response.IsSuccessful {
		routeHelper.SetSuccessResponse(ginContext, &model.ResponseData{Data: response})
	} else {
		routeHelper.SetFailureResponseWithStatusCode(ginContext, http.StatusFailedDependency, response.Error)
	}
}
