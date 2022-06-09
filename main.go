package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zawachte/bimah/internal/api"
	"github.com/zawachte/bimah/internal/providers"
)

func main() {

	r := gin.Default()
	srv, err := providers.NewProvider(context.Background(), providers.ProviderParams{
		DatabaseUrl: "", //"postgres://zach:zach@localhost:5400/zach",
	})

	if err != nil {
		panic(err)
	}
	//	caas.RegisterHandlers(r, srv)
	api.RegisterHandlersWithOptions(r, srv, api.GinServerOptions{BaseURL: "v1"})
	r.Run()
}
