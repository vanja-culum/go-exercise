package main

import (
	"fmt"
	"go-exercise/ds"
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
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Logger)

	r.Get("/", rootHandler.HandleGetRoot)

	r.Get("/dsa", dsaHandler.HandleGetDSAFiles)

	r.Post("/dsa/bst/generate", dsaHandler.HandlePostBstGenerate)

	t := ds.BST[int]{}

	t.Insert(10)
	t.Insert(5)

	t.Insert(15)

	t.Insert(30)
	// t.Insert(40)

	// fmt.Println(t.InOrder())
	fmt.Println(t.GenerateLinks())

	links := t.GenerateLinks()

	for _, link := range links {

		if link.Parent != nil {
			fmt.Println("link parent val", link.Parent.Val)
		}
		
		if link.Child != nil {
			fmt.Println("link child val", link.Child.Val)
		}
	}

	// for _, node := range t.InOrderNode() {
	// 	level, _ := t.FindLevel(node.Val)
	// 	fmt.Println(node.Val, level)
	// }
	http.ListenAndServe(":8080", r)
}
