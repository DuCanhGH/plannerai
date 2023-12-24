package middleware

import "net/http"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get("Access-Control-Allow-Origin") == "" {
			http.Error(w, "", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
