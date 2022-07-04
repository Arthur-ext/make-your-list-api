package router

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(10 * time.Second),
	)

	router.Route("/api", func(r chi.Router) {
		
	})

	return router
}