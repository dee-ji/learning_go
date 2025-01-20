package main

import (
	"context"
	"fmt"
	"net/http"
)

// Level is a custom type with string as the underlying type
type Level string

// Define constants for the log levels
const (
	Debug Level = "debug"
	Info  Level = "info"
)

// contextKey is a type to prevent collisions in context keys
type contextKey string

const logLevelKey contextKey = "logLevel"

// StoreLogLevelInContext stores the log level in the given context
func StoreLogLevelInContext(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, logLevelKey, level)
}

// GetLogLevelFromContext extracts the log level from the context
// If the log level is not set or invalid, it returns an empty string.
func GetLogLevelFromContext(ctx context.Context) Level {
	if val, ok := ctx.Value(logLevelKey).(Level); ok && (val == Debug || val == Info) {
		return val
	}
	return ""
}

// Log function logs messages at the appropriate log level
func Log(ctx context.Context, level Level, message string) {
	inLevel := GetLogLevelFromContext(ctx) // Extract log level from context

	// Determine if the message should be logged based on the level
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Debug || inLevel == Info) {
		fmt.Println(message)
	}
}

// LogLevelMiddleware extracts the log level from the query parameter and store it in the context
func LogLevelMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logLevel := r.URL.Query().Get("log_level")
		var level Level

		// Validate the log level
		if logLevel == string(Debug) {
			level = Debug
		} else if logLevel == string(Info) {
			level = Info
		}

		// Store the log level in the context if valid
		ctx := StoreLogLevelInContext(r.Context(), level)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	// Example HTTP handler to demonstrate logging
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log(r.Context(), Debug, "This is a debug message")
		Log(r.Context(), Info, "This is an info message")
		fmt.Fprintln(w, "Check logs for messages")
	})

	// Wrap the handler with the log level middleware
	http.Handle("/", LogLevelMiddleware(handler))

	// Start the HTTP server
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
