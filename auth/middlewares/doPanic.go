package middlewares

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// this middleware will deliberately pass error to next
func DoPanic(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc) {
	fmt.Println("Panic!")
	next(errors.New("uh oh"))
}
