package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    Input
		wantErr bool
	}{
		{
			name:  "Valid input",
			input: []byte("id123\n+\n10\n20\n"),
			want: Input{
				Id:   "id123",
				Op:   "+",
				Val1: 10,
				Val2: 20,
			},
			wantErr: false,
		},
		{
			name:    "Invalid integer value",
			input:   []byte("id123\n+\n10\ninvalid\n"),
			want:    Input{},
			wantErr: true,
		},
		{
			name:    "Missing lines",
			input:   []byte("id123\n+\n10\n"),
			want:    Input{},
			wantErr: true,
		},
		{
			name:    "Empty input",
			input:   []byte(""),
			want:    Input{},
			wantErr: true,
		},
		{
			name:    "Only ID provided",
			input:   []byte("id123\n"),
			want:    Input{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	in := make(chan []byte, 1)
	out := make(chan Result, 1)

	// Send valid data
	in <- []byte("id123\n+\n10\n20\n")
	close(in)

	go DataProcessor(in, out)

	result := <-out
	expected := Result{Id: "id123", Value: 30}

	if result != expected {
		t.Errorf("DataProcessor() = %v, want %v", result, expected)
	}
}

func TestWriteData(t *testing.T) {
	results := make(chan Result, 1)
	results <- Result{Id: "id123", Value: 30}
	close(results)

	var buf bytes.Buffer
	WriteData(results, &buf)

	expected := "id123:30\n"
	if buf.String() != expected {
		t.Errorf("WriteData() = %v, want %v", buf.String(), expected)
	}
}

func TestNewController(t *testing.T) {
	ch := make(chan []byte, 1)
	handler := NewController(ch)

	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantBody   string
	}{
		{
			name:       "Valid request",
			body:       "id123\n+\n10\n20\n",
			wantStatus: http.StatusAccepted,
			wantBody:   "OK: 1",
		},
		{
			name:       "Channel full",
			body:       "id123\n+\n10\n20\n",
			wantStatus: http.StatusServiceUnavailable,
			wantBody:   "Too Busy: 1",
		},
		{
			name:       "Invalid body",
			body:       "",
			wantStatus: http.StatusBadRequest,
			wantBody:   "Bad Input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			rec := httptest.NewRecorder()

			handler.ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status code = %v, want %v", rec.Code, tt.wantStatus)
			}
			if !strings.Contains(rec.Body.String(), tt.wantBody) {
				t.Errorf("body = %q, want %q", rec.Body.String(), tt.wantBody)
			}
		})
	}
}

func TestMainIntegration(t *testing.T) {
	ch1 := make(chan []byte, 1)
	ch2 := make(chan Result, 1)

	var buf bytes.Buffer
	var mu sync.Mutex // Protects access to buf

	// Start the DataProcessor and WriteData goroutines
	go DataProcessor(ch1, ch2)
	go func() {
		for r := range ch2 {
			mu.Lock()
			buf.Write([]byte(fmt.Sprintf("%s:%d\n", r.Id, r.Value)))
			mu.Unlock()
		}
	}()

	// Send data to the input channel
	ch1 <- []byte("id123\n+\n10\n20\n")
	close(ch1)

	// Wait for the goroutines to complete processing
	time.Sleep(100 * time.Millisecond)

	// Access the buffer safely
	mu.Lock()
	result := buf.String()
	mu.Unlock()

	expected := "id123:30\n"
	if result != expected {
		t.Errorf("Integration test = %q, want %q", result, expected)
	}
}

func FuzzParser(f *testing.F) {
	// Add some seed inputs for the fuzzer
	f.Add([]byte("id123\n+\n10\n20\n"))        // Valid input
	f.Add([]byte("id123\n+\n10\ninvalid\n"))   // Invalid number
	f.Add([]byte("id123\n/\n100\n0\n"))        // Division by zero
	f.Add([]byte(""))                          // Empty input
	f.Add([]byte("id123\n*\n\n\n"))            // Missing numbers
	f.Add([]byte("id123\n+\n10\n20\nExtra\n")) // Extra data

	f.Fuzz(func(t *testing.T, data []byte) {
		_, err := parser(data)
		if err != nil {
			// Ensure we don't panic; error is expected for invalid inputs
			t.Logf("Expected error: %v", err)
		}
	})
}
