package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// TimeoutMiddleware creates a middleware that sets a context with a timeout for each request.
// timeoutMs specifies the timeout duration in milliseconds.
func TimeoutMiddleware(timeoutMs int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Convert milliseconds to a time.Duration
			timeoutDuration := time.Duration(timeoutMs) * time.Millisecond

			// Create a context with a timeout
			ctx, cancel := context.WithTimeout(r.Context(), timeoutDuration)
			defer cancel() // Ensure the cancel function is called to release resources

			// Attach the new context with the timeout to the request
			r = r.WithContext(ctx)

			// Call the next handler in the chain
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	// Example handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
			http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			return
		case <-time.After(1 * time.Second): // Simulate processing time
			fmt.Fprintln(w, "Request processed successfully")
		}
	})

	// Create middleware with a 2-second timeout
	timeoutMiddleware := TimeoutMiddleware(2000)

	// Wrap the handler with the middleware
	http.Handle("/", timeoutMiddleware(handler))

	// Start the server
	http.ListenAndServe(":8080", nil)
}
