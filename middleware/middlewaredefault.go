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
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Set-Cookie")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.Write([]byte("allowed"))
			return
		}

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
