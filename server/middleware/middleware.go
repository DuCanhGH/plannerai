package middleware

import (
	"net/http"
	"os"
	"strings"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Referer"), os.Getenv("CLIENT_URL")) && os.Getenv("CLIENT_URL") != "*" {
			http.Error(w, "", http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
