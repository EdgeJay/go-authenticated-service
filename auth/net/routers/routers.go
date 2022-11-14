package routers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, _ interface{}) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Server error")
	}
	return router
}
