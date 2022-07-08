package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", listGifts)
	r.Post("/", addGift)
	r.Patch("/", updateGift)

	return r
}

func listGifts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List wedding gifts!"))
}

func addGift(w http.ResponseWriter, r *http.Request) {

}

func updateGift(w http.ResponseWriter, r *http.Request) {
	
}