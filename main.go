package main


import (
	"github.com/nonCriticInc/heimdall/config"
	"github.com/nonCriticInc/heimdall/pkg/routes"
)

func main() {
	srv := config.New()
	routes.Routes(srv)
	srv.Logger.Fatal(srv.Start(":8081"))
}
