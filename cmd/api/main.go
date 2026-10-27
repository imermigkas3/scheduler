package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
) 
func main() {
	mux := http.NewServeMux()

	// if the function takes an interface then you can pass either a value that implemnets that or a pointer to a value taht implments that
	// https://go.dev/doc/faq#pointer_to_interface

	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
    // request against a list of registered patterns and calls the handler for the
    // pattern that most closely matches the URL.


	// need {$} because GET / means get any GET whose path starts with / its a prefix match
	// {$} means the path ends here
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
 
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"}) // map literal
	})

	log.Println("listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}