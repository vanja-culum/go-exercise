package main

import (
	"go-exercise/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()

	rootHandler := handler.NewRootHandler()
	codeHandler := handler.NewCodeHandler()

	r.Use(middleware.Logger)
 
	r.Get("/", rootHandler.HandleGetRoot)

	r.Get("/dsa", codeHandler.HandleGetDSAFiles)

	http.ListenAndServe(":3000", r)
}