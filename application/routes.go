package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thanh2k4/Chat-app/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})
	router.Route("/users", loadUserRoutes)
	return router
}

func loadUserRoutes(router chi.Router) {
	userHandler := &handler.User{}
	router.Get("/", userHandler.List)
	router.Post("/", userHandler.Create)
	router.Get("/{id}", userHandler.GetByID)
	router.Put("/{id}", userHandler.UpdateByID)
	router.Delete("/{id}", userHandler.DeleteByID)
}
