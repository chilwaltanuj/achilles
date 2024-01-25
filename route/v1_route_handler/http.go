package v1RouteHandler

import (
	"achilles/constant"
	"achilles/helper"
	"achilles/model"
	"context"
	"net/http"

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

	contextWithTimeout, cancel := context.WithTimeout(ginContext, helper.ApplicationConfiguration().Config.Timmeout)
	defer cancel()

	if response := helper.Execute[model.Joke_Random](contextWithTimeout, request); response.IsSuccessful {
		routeHelper.SetSuccessResponse(ginContext, &model.ResponseData{Data: response})
	} else {
		routeHelper.SetFailureResponseWithStatusCode(ginContext, http.StatusFailedDependency, response.Error)
	}
}
