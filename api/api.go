package api

import (
	"wedding_gifts/api/router"
	"wedding_gifts/internal/app"

	"github.com/go-chi/chi/v5"
)

type API struct {
	router *chi.Mux
}

func newAPI() *API {
	application := app.NewAPP()

	router := router.InitRouter(application)

	return &API{
		router: router,
	}
}
