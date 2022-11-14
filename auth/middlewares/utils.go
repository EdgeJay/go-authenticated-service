package middlewares

import (
	"fmt"
	"net/http"

	"github.com/EdgeJay/go-authenticated-service/auth/utils/arrays"
	"github.com/julienschmidt/httprouter"
)

type NextFunc func(err any)

type MiddlewareFunc func(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc)

func NewWrappedHandler(handlerToWrap httprouter.Handle, middlewares ...MiddlewareFunc) httprouter.Handle {
	var next NextFunc

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		middlewares = append(middlewares, func(w http.ResponseWriter, r *http.Request, params httprouter.Params, next NextFunc) {
			handlerToWrap(w, r, params)
		})

		it := arrays.NewIterator(middlewares)

		next = func(err any) {
			if _, ok := err.(error); ok {
				// err is error object, need to throw error
				fmt.Printf("Error detected: %v\n", err)
				panic(err)
			}

			if m, done := it.NextFunc(); !done {
				m(w, r, params, next)
			}
		}
		next(nil)
	}
}
