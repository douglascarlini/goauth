package main

import (
    "os"
    "fmt"
    "context"
    "net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleIndex)

	server := &http.Server{ Addr: ":8080", Handler: cors(mux) }

	go func() {
		fmt.Println("Listening on http://localhost:8080/")
		if err := server.ListenAndServe(); err != nil {
			if err.Error() != "http: Server closed" {
				fmt.Printf("Error: %s\n", err)
			}
		}
	}()

	<-stopSignal()

	fmt.Println("Shutting down the server...")
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("Error shutting down the server: %s\n", err)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	response := os.Getenv("OAUTH_RESPONSE")
	token := r.URL.Query().Get("token")
	fmt.Printf("TOKEN: %s\n", token)
	fmt.Fprintf(w, response)
	sendStopSignal()
}

var stopSignalChan = make(chan struct{}, 1)

func stopSignal() <-chan struct{} {
	return stopSignalChan
}

func sendStopSignal() {
	close(stopSignalChan)
}

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}