package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/middlewares"
	"github.com/EdgeJay/go-authenticated-service/auth/net/routers"
	"github.com/EdgeJay/go-authenticated-service/auth/net/routes"
	"github.com/EdgeJay/go-authenticated-service/auth/utils/env"
)

func main() {
	env.LoadEnv()

	router := routers.New()
	router.GET("/", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger))
	router.GET("/interrupt", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger, middlewares.Interrupt))
	router.GET("/error", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger, middlewares.DoPanic))

	fmt.Println(env.GetPort())
	log.Fatal(http.ListenAndServe(env.GetPort(), router))
}
