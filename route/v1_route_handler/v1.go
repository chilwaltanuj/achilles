package v1RouteHandler

import (
	routeHelper "achilles/route/helper"
	middlewareHandler "achilles/route/middleware_handler"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

// Init initializes the version 1 routes.
func Init(router *gin.Engine) {
	ginEngine = router
}

// AddRouteHandlers adds route handlers to the version 1 group.
func AddRouteHandlers() error {
	versionOne := ginEngine.Group("v1")
	AttachMiddlewares(versionOne)

	versionOne.GET("test/crash", Crash)
	versionOne.GET("test/success", Success)

	return nil
}

// AttachMiddlewares attaches middlewares to the specified group.
func AttachMiddlewares(group *gin.RouterGroup) {
	group.Use(middlewareHandler.Recovery())
	group.Use(PreRequestMiddlewares())
	group.Use(PostRequestMiddlewares())

	ginEngine.HandleMethodNotAllowed = true
	ginEngine.NoMethod(middlewareHandler.MethodNotSupported)
	ginEngine.NoRoute(middlewareHandler.RouteNotSupported)
}

// PreRequestMiddlewares returns a middleware function for pre-request processing.
func PreRequestMiddlewares() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		routeHelper.BuildAndSetRequestMetaInContext(ginContext)
	}
}

// PostRequestMiddlewares returns a middleware function for post-request processing.
func PostRequestMiddlewares() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.Next()
		routeHelper.UpdateRequestMetaDataInContext(ginContext)
		middlewareHandler.RenderResponse(ginContext)
	}
}
