package api

import (
	"wedding_gifts/api/router"
	"wedding_gifts/internal/app"

	"github.com/go-chi/chi/v5"
)

type API struct {
	application *app.APP
	router *chi.Mux
}

func newAPI() *API {
	application := app.NewAPP()

	router := router.InitRouter()

	return &API{
		application: application,
		router: router,
	}
}