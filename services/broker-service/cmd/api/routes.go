package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// In order to share the routes in the project, add a reciever for it.
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// ** Specify who is allowed to connect ** //
	// Use appends a middleware handler to the Mux middleware stack. -> cors indicates a middleware here
	// func (mx *Mux) With(middlewares ...func(http.Handler) http.Handler) Router
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// A simple middleware to check the service is alive or not
	mux.Use(middleware.Heartbeat("/ping"))

	mux.Post("/", app.Broker)
	mux.Post("/handle", app.HandleSubmission)

	return mux
}
