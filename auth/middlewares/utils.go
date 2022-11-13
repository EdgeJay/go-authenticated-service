package middlewares

import (
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/utils/arrays"
	"github.com/julienschmidt/httprouter"
)

type NextFunc func()

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc)

func NewWrappedHandler(handlerToWrap httprouter.Handle, middlewares ...MiddlewareFunc) httprouter.Handle {
	var next NextFunc

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		middlewares = append(middlewares, func(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc) {
			handlerToWrap(w, r, params)
		})

		it := arrays.NewIterator(middlewares)

		next = func() {
			if m, done := it.NextFunc(); !done {
				m(w, r, params, next)
			}
		}
		next()
	}
}
