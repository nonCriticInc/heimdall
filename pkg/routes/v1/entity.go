package v1

import (

	"github.com/labstack/echo"
	v1 "github.com/nonCriticInc/heimdall/pkg/v1"

)

func EntityRouter(g *echo.Group) {

	g.POST("", v1.Create)
	g.GET("/:id/", v1.FindById)
}
