package middleware

import (
	"net/http"
)

type DefaultMiddleware struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (dm *DefaultMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(w, r)
	})
}

func (dm *DefaultMiddleware) RegisterMiddlewareDefault(next func(next http.Handler) http.Handler) {
	dm.middleware = append(dm.middleware, next)
}

func (dm *DefaultMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &dm.ServeMux
	for _, next := range dm.middleware {

		current = next(current)
	}

	current.ServeHTTP(w, r)
}
