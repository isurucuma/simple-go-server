package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

// formHandler will handle the form submission
// It will parse the form and then print the values
// here ResponseWriter is used to write the response, when the function ends the response will be sent to the client
// Request is used to get the request from the client
// here r is a pointer to the Request struct
func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	// ParseForm parses the raw query from the URL and updates r.Form
	// It will also parse the body of the request
	// calling this function is necessary to get the form values
	r.ParseForm()
	fmt.Fprintf(w, "Name = %s\n", r.FormValue("fname"))
	fmt.Fprintf(w, "Address = %s\n", r.FormValue("lname"))
}

func main() {
	// FileServer will going to serve all the static files in the static folder
	// It will simply return the files with the filename as the path
	// FileServer is an http.Handler
	FileServer := http.FileServer(http.Dir("./static"))

	// Handle will register the handler for the given pattern
	http.Handle("/", FileServer)

	// HandleFunc will register the handler function for the given pattern
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is listening...")

	// ListenAndServe will listen on the TCP network address and then calls Serve with handler to handle requests on incoming connections
	// when the handler is nil, DefaultServeMux is used. that means all the above registered handlers will be used
	if err := http.ListenAndServe(":8181", nil); err != nil {
		log.Panic(err)
	}
}
