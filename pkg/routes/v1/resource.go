package v1

import (
	"github.com/labstack/echo"
	v1 "github.com/nonCriticInc/heimdall/pkg/v1"
)

func ResourceRouter(g *echo.Group) {

	g.POST("", v1.CreateReources)
	g.GET("/:id", v1.FindResourceById)

}

