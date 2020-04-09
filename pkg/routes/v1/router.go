package v1

import (
	"github.com/labstack/echo"
)

func V1Router(g *echo.Group) {
	entityMonitor := g.Group("/entities")
	EntityRouter(entityMonitor)

	organizationMonitor:= g.Group("/organizations")
	OrganizationRouter(organizationMonitor)


	applicationMonitor:= g.Group("/applications")
	ApplicationRouter(applicationMonitor)

	resourceMonitor:= g.Group("/resources")
	ResourceRouter(resourceMonitor)


}
