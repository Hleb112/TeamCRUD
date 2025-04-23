package server

import (
	"log"
	"net/http"
	"time"
)

func CreateServer() {
	myServer := &http.Server{
		// set the server address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// launch the server
	log.Fatal(myServer.ListenAndServe())
}
