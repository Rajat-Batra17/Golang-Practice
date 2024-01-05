package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello User! Have a nice day ahead")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parseform() error %v\n", err)
		return
	}
	fmt.Fprintf(w, "Welcome to Rajat's Form\n")
	fmt.Fprintf(w, "POST request successfully made\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name := %v\n", name)
	fmt.Fprintf(w, "Address := %v\n", address)
}
func main() {

	server := http.FileServer(http.Dir("./static"))
	// server variable holds an HTTP handler that serves files from the "static" directory.

	http.Handle("/", server) // will route it to index.html in static folder

	/*
		http.HandleFunc() : associates a function with a specific URL pattern,
		enabling you to define how HTTP requests with that pattern should be handled.
	*/
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler) // tell http to call homehandler function
	// when particular route is called

	fmt.Println("Starting server at PORT 8000")
	// starts an HTTP server and listens for incoming requests on the specified address and port,

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Error", err)
	}

}
