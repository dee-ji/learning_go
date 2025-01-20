package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"

	//"fmt"
	"net/http"
	"time"
)

type TimeHandler struct{}

func (t *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	if r.Header.Get("Accept") == "application/json" {
		w.Header().Set("Content-Type", "application/json")
		timeData := map[string]interface{}{
			"day_of_week":  now.Weekday().String(),
			"day_of_month": now.Day(),
			"month":        now.Month().String(),
			"year":         now.Year(),
			"hour":         now.Hour(),
			"minute":       now.Minute(),
			"second":       now.Second(),
		}

		data, _ := json.MarshalIndent(timeData, "", "  ")
		w.Write(data)
	} else {
		w.Write([]byte(now.String()))
	}
}

func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		dur := time.Since(start)

		options := &slog.HandlerOptions{Level: slog.LevelDebug}
		handler := slog.NewJSONHandler(os.Stderr, options)
		mySlog := slog.New(handler)
		mySlog.Info("request time",
			"path", r.URL.Path,
			"duration", dur,
		)
	})
}

func main() {
	mux := http.NewServeMux()

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	mux.Handle("/time", RequestTimer(&TimeHandler{}))
	err := s.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
