package route

import (
	"achilles/helper"
	v1RouteHandler "achilles/route/v1_route_handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the application's routes.
func SetupRouter() (*gin.Engine, error) {
	ginEngine := gin.New()

	v1RouteHandler.Init(ginEngine)
	helper.GetGlobalLogger().Info("router up")
	v1RouteHandler.AddRouteHandlers()

	return ginEngine, nil
}
