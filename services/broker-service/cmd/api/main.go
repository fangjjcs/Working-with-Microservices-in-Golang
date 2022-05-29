package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct {
}

func main() {

	// Declare the app
	app := Config{}

	log.Printf("Starting Broker Service on Port %s", PORT)

	// Define http server
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%s", PORT),
		Handler: app.routes(),
	}

	// Start the http server
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
