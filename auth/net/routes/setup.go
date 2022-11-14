package routes

import (
	"github.com/EdgeJay/go-authenticated-service/auth/middlewares"
	"github.com/julienschmidt/httprouter"
)

func SetupRoutes(router *httprouter.Router) {
	router.GET("/", middlewares.NewWrappedHandler(Index, middlewares.Logger))
	router.GET("/interrupt", middlewares.NewWrappedHandler(Index, middlewares.Logger, middlewares.Interrupt))
	router.GET("/error", middlewares.NewWrappedHandler(Index, middlewares.Logger, middlewares.DoPanic))
}
