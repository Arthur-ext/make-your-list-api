package router

import (
	"time"

	"wedding_gifts/api/handler"
	"wedding_gifts/internal/app"
	apiMiddleware "wedding_gifts/api/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter(a app.APP) *chi.Mux {
	router := chi.NewRouter()
	
	router.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.Timeout(10 * time.Second),
		apiMiddleware.SetHeaderContentType("application/json"),
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/wedding-gifts", handler.NewGiftHandler(a.Gifts).Routes())
	})

	return router
}