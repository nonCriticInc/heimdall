package routes
import (

	"github.com/labstack/echo"
	v1 "github.com/nonCriticInc/heimdall/pkg/routes/v1"
)

func Routes(e *echo.Echo) {
	v1Monitor := e.Group("/api/v1")
	v1.V1Router(v1Monitor)

}
