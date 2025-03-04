package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", handler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
