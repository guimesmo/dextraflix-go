package routes

import (
	"github.com/guimesmo/dextraflix-go/controllers"
	"github.com/labstack/echo"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
}
