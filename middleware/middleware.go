package middleware

import (
	"net/http"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

type GoMiddleware struct {
}

func InitMidleware() *GoMiddleware {
	return &GoMiddleware{}
}

func (gm *GoMiddleware) Method(m string) Adapter {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if m != r.Method {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}
			f(w, r)
		}
	}
}

func (gm *GoMiddleware) ApplyMiddleware(f http.HandlerFunc, mw ...Adapter) http.HandlerFunc {
	for _, m := range mw {
		f = m(f)
	}

	return f
}
