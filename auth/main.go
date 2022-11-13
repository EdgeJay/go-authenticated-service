package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/middlewares"
	"github.com/julienschmidt/httprouter"
)

func logger(w http.ResponseWriter, r *http.Request, params httprouter.Params, next middlewares.NextFunc) {
	fmt.Println("Logging before")
	next()
	fmt.Println("Logging after")
}

// this middleware does not call next function
func interrupt(w http.ResponseWriter, r *http.Request, params httprouter.Params, _ middlewares.NextFunc) {
	fmt.Println("Interrupt execution of next function")
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Output content")
	fmt.Fprint(w, "Hello")
}

func main() {
	router := httprouter.New()
	router.GET("/", middlewares.NewWrappedHandler(index, logger))
	router.GET("/interrupt", middlewares.NewWrappedHandler(index, logger, interrupt))

	log.Fatal(http.ListenAndServe(":8080", router))
}
