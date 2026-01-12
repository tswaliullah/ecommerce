package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			next.ServeHTTP(w, r)

			diff := time.Since(start)
//   log the method, path, and time taken to process the request
			log.Println(r.Method, r.URL.Path, diff)
		})
}
