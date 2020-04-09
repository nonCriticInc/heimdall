package v1

import (
	"github.com/labstack/echo"
	v1 "github.com/nonCriticInc/heimdall/pkg/v1"
)

func ApplicationRouter(g *echo.Group) {

	g.POST("", v1.CreateApplications)
	g.GET("/:id", v1.FindApplicationById)

}
