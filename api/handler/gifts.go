package handler

import (
	"encoding/json"

	"net/http"

	"wedding_gifts/internal/app/controller"
	"wedding_gifts/internal/model"

	"github.com/go-chi/chi/v5"
)

type GiftHandler struct {
	controller controller.GiftsController
}

func NewGiftHandler(c controller.GiftsController) GiftHandler {
	return GiftHandler{
		controller: c,
	}
}

func (h GiftHandler) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", h.addGift)
	r.Get("/", h.listGifts)
	r.Get("/{gift-id}", h.getGift)
	r.Patch("/{gift-id}", h.updateGift)
	r.Delete("/{gift-id}", h.deleteGift)

	return r
}

func (h GiftHandler) addGift(w http.ResponseWriter, r *http.Request) { 
	var gift model.Gift
	if err := json.NewDecoder(r.Body).Decode(&gift); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	
	if err := h.controller.Create(gift); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h GiftHandler) listGifts(w http.ResponseWriter, r *http.Request) {
	gifts, err := h.controller.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	giftsJson, err := json.Marshal(gifts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(giftsJson)
}

func (h GiftHandler) getGift(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "gift-id")

	gift, err := h.controller.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	giftJson, err := json.Marshal(gift)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(giftJson)
}

func (h GiftHandler) updateGift(w http.ResponseWriter, r *http.Request) {
	var payload map[string]string

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	payload["id"] = chi.URLParam(r, "gift-id")
	
	if err := h.controller.UpdateAssigner(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h GiftHandler) deleteGift(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "gift-id")
	
	if err := h.controller.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}