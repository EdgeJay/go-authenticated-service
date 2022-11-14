package middlewares

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// this middleware does not call next function
func Interrupt(w http.ResponseWriter, r *http.Request, params httprouter.Params, _ NextFunc) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Interrupted call")
}
