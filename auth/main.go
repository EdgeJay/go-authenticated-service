package main

import (
	"log"
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/middlewares"
	"github.com/EdgeJay/go-authenticated-service/auth/net/routers"
	"github.com/EdgeJay/go-authenticated-service/auth/net/routes"
)

func main() {
	router := routers.New()
	router.GET("/", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger))
	router.GET("/interrupt", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger, middlewares.Interrupt))
	router.GET("/error", middlewares.NewWrappedHandler(routes.Index, middlewares.Logger, middlewares.DoPanic))

	log.Fatal(http.ListenAndServe(":8080", router))
}
