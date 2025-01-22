package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Result struct {
	Id    string
	Value int
}

type Input struct {
	Id   string
	Op   string
	Val1 int
	Val2 int
}

func parser(data []byte) (Input, error) {
	// Parse the data
	lines := bytes.Split(data, []byte("\n"))
	// Validate the number of lines
	if len(lines) < 4 {
		return Input{}, fmt.Errorf("invalid input: not enough lines")
	}
	// Parse ID and operator
	id := string(lines[0])
	op := string(lines[1])
	// Parse the two integer values
	val1, err := strconv.Atoi(string(lines[2]))
	if err != nil {
		return Input{}, fmt.Errorf("invalid value for val1: %w", err)
	}
	val2, err := strconv.Atoi(string(lines[3]))
	if err != nil {
		return Input{}, fmt.Errorf("invalid value for val2: %w", err)
	}
	// Return the parsed input
	return Input{
		Id:   id,
		Op:   op,
		Val1: val1,
		Val2: val2,
	}, nil
}

func DataProcessor(in <-chan []byte, out chan<- Result) {
	for data := range in {
		input, err := parser(data)
		if err != nil {
			continue
		}
		var calc int
		switch input.Op {
		case "+":
			calc = input.Val1 + input.Val2
		case "-":
			calc = input.Val1 - input.Val2
		case "*":
			calc = input.Val1 * input.Val2
		case "/":
			if input.Val2 != 0 {
				calc = input.Val1 / input.Val2
			} else {
				continue
			}
		default:
			continue
		}
		// sum numbers in the data
		// write to another channel
		result := Result{
			Id:    input.Id,
			Value: calc,
		}
		out <- result
	}
	close(out)
}

func WriteData(in <-chan Result, w io.Writer) {
	for r := range in {
		// write the output data to writer
		// each line is id:result
		w.Write([]byte(fmt.Sprintf("%s:%d\n", r.Id, r.Value)))
	}
}

func NewController(out chan []byte) http.Handler {
	var numSent int
	var numRejected int
	var mu sync.Mutex // Mutex to protect shared variables

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil || len(data) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Bad Input"))
			return
		}

		select {
		case out <- data:
			mu.Lock()
			numSent++
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte("OK: " + strconv.Itoa(numSent)))
			mu.Unlock()
		default:
			mu.Lock()
			numRejected++
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("Too Busy: " + strconv.Itoa(numRejected)))
			mu.Unlock()
		}
	})
}

func main() {
	// set everything up
	ch1 := make(chan []byte, 100)
	ch2 := make(chan Result, 100)
	go DataProcessor(ch1, ch2)
	f, err := os.Create("results.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	go WriteData(ch2, f)
	err = http.ListenAndServe(":8080", NewController(ch1))
	if err != nil {
		fmt.Println(err)
	}
}
