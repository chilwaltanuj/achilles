package route

import (
	"fmt"

	v1RouteHandler "achilles/route/v1_route_handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

// SetupRouter initializes and configures the application's routes.
func SetupRouter(container *dig.Container) (*gin.Engine, error) {
	var router *gin.Engine
	router = gin.New()
	if err := container.Invoke(v1RouteHandler.Init); err != nil {
		fmt.Println(err)
	}
	v1RouteHandler.AddRouteHandlers(router)

	return router, nil
}
