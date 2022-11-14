package main

import (
	"log"
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/net/routers"
	"github.com/EdgeJay/go-authenticated-service/auth/net/routes"
	"github.com/EdgeJay/go-authenticated-service/auth/utils/env"
	"github.com/julienschmidt/httprouter"
)

type wrappedRouter struct {
	router *httprouter.Router
}

func (r *wrappedRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)

}

func startServer(router *httprouter.Router) {
	log.Println("Listening to port", env.GetPort())
	log.Fatal(http.ListenAndServe(env.GetPort(), &wrappedRouter{router}))
}

func main() {
	env.LoadEnv()
	router := routers.New()
	routes.SetupRoutes(router)
	startServer(router)
}
