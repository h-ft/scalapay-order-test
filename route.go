package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// A function to handle the request and pass them to its corresponding router
func HandleRequests() {
	// Creates a new instance of a mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	// Replace http.HandleFunc with myRouter.HandleFunc
	muxRouter.HandleFunc("/", homePage)
	muxRouter.HandleFunc("/order", createOrder).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", muxRouter))
}
