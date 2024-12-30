package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[Start] %s", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("[End] %s", r.URL)
	})
}
