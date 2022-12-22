package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func routes() http.Handler {
	router := chi.NewRouter

	// specify who is allowed to connect
	router.Use(cors.Handler)
}
