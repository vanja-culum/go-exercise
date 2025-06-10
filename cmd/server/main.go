package main

import (
	"go-exercise/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	r := chi.NewRouter()

	rootHandler := handler.NewRootHandler()
	dsaHandler := handler.NewDsaHandler()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"localhost:3000", "https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Establishment-ID"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)

	r.Get("/", rootHandler.HandleGetRoot)

	r.Get("/dsa", dsaHandler.HandleGetDSAFiles)

	r.Post("/dsa/bst/generate", dsaHandler.HandlePostBstGenerate)

	http.ListenAndServe(":8080", r)
}
