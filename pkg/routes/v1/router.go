package v1

import (
	"github.com/labstack/echo"
)

func V1Router(g *echo.Group) {
	entityMonitor := g.Group("/entities")
	EntityRouter(entityMonitor)
}
