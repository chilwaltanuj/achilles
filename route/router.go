package route

import (
	v1RouteHandler "achilles/route/v1_route_handler"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures the application's routes.
func SetupRouter() (*gin.Engine, error) {
	ginEngine := gin.New()
	v1RouteHandler.Init(ginEngine)
	v1RouteHandler.AddRouteHandlers()

	return ginEngine, nil
}
