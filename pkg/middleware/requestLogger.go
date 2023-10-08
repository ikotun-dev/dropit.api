package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Record the start time of the request
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Calculate the duration of the request
		duration := time.Since(startTime)

		// Log the request details
		fmt.Printf(
			"[%s] %s %s %s %s %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			r.Proto,
			duration,
		)
	})
}
