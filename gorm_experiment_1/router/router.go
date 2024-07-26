package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gormex/handler"
	"net/http"
)

func Router() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	//	route
	router.Get("/notes/", handler.AllNotes)
	router.Post("/notes/", handler.CreateNote)
	router.Put("/notes/{id}", handler.UpdateNote)
	router.Delete("/notes/{id}", handler.DeleteNote)

	return router
}
