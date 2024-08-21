package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"test-task/handler"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Post("/operation", handler.OperationHandler)
	})

	return router
}
