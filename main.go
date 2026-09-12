package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type CalcRequest struct {
	A 	int    `json:"a"`
	B 	int	   `json:"b"`
	Op 	string `json:"op"`
}

type CalcResponse struct {
	Result  int    `json:"result"`
	Error   string `json:"error,omitempty"`
}

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "add":
		return a + b, nil
	case "sub":
		return a - b, nil
	case "mul":
		return a * b, nil
	case "div":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "только POST", http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "плохой JSON", http.StatusBadRequest)
		return
	}

	resultCh := make(chan int)
	errCh := make(chan error)

	go func() {
		res, err := calculate(req.A, req.B, req.Op)
		if err != nil {
			errCh <- err
			return
		}
		resultCh <- res
	}()

	var resp CalcResponse

	select {
	case res := <-resultCh:
		resp.Result = res
	case err := <-errCh:
		resp.Error = err.Error()
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("-> %s %s\n", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Printf("<- %s %s (%v)\n", r.Method, r.URL.Path, time.Since(start))
	}) 
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/calc", calcHandler)

	handler := loggingMiddleware(mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatalf("ошибка %v", err)
	}
}