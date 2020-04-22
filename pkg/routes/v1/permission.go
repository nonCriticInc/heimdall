package v1

import (
	v1 "github.com/nonCriticInc/heimdall/pkg/v1"
	"github.com/labstack/echo"
)

func PermissionRouter(g *echo.Group) {

	g.POST("", v1.CreatePermissions)
	g.GET("/:id", v1.FindPermissionById)
	g.GET("/:id/childs", v1.FindChildPermissions)

}

