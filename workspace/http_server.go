package main

import (
	"log"
	"net/http"
)

func helloGoHandler(
	w http.ResponseWriter, /* 	http.ResponseWriter: Construct the HTTP response 	*/
	r *http.Request) { /* 		http.Request: Retrieve the request 					*/
	w.Write([]byte("Hello net/http!\n"))
}

func submain() {
	http.HandleFunc("/", helloGoHandler) /* Register helloGoHandler as the handler function for any request that matches a given pattern: "/" is the root pattern */

	log.Fatal(http.ListenAndServe(":8080", nil))
	/*
		http.ListenAndServe(addr string, handler http.Handler)
			ListenAndServer always stops the execution flow, only returning in the vent of an error
							always returns a no-nil error, so it is important to log it
	*/
}

/*
ListenAndServe
	Starts an HTTP server with:
	- given address
	- a handler: if handler is nill, usually in net/http lib, DefaultServeMux value is used

Handler:
	A Handler is any type that satisfies the Handler interface by providing a ServeHTTP Method
	- type Handler interface { ServeHTTP(ResponseWriter, *Request)}

Handler acts as a multiplexer. The mux compares the requested URL to the registered patterns and call the handler function
DefaultServeMux is a global value of type serveMux which implements default HTTP multiplexer logic
*/
