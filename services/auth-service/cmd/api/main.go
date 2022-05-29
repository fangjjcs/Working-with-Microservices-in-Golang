package main

import (
	"database/sql"
	"fangjjcs/auth/data"
	"fmt"
	"log"
	"net/http"
)

const PORT = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	app := Config{}
	log.Println("Starting authentication service")

	// TODO connect DB

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
