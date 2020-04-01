package main


import (
	"github.com/nonCriticInc/heimdall/config"

)

func main() {
	srv := config.New()
	srv.Logger.Fatal(srv.Start(":8081"))
}
