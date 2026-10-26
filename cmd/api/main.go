package main

import (
	"fmt"
	"log"
	"net/http"
) 
func main() {
	mux := http.NewServeMux()

	// if the function takes an interface then you can pass either a value that implemnets that or a pointer to a value taht implments that
	// https://go.dev/doc/faq#pointer_to_interface
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	log.Println("listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}