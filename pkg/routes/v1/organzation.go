package v1

import (
	"github.com/labstack/echo"
	v1 "github.com/nonCriticInc/heimdall/pkg/v1"
)

func OrganizationRouter(g *echo.Group) {

	g.POST("", v1.CreateOrganizations)
	g.GET("/:id", v1.FindOrganizationById)
	g.GET("/:id/applications", v1.FindApplicationsByOrganization)

}
