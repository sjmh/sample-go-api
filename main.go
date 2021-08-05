package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sjmh/testapi/items"
)

type HTTPService struct {
	itemRepository items.Repository
}

func (h *HTTPService) getItemById(w http.ResponseWriter, r *http.Request) {
	itemId := chi.URLParam(r, "id")
	item, _ := h.itemRepository.GetItemByID(r.Context(), itemId)
	json.NewEncoder(w).Encode(item)
}

func (h *HTTPService) getItems(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.itemRepository.GetItems(r.Context()))
}

func main() {
	r := chi.NewRouter()
	handler := HTTPService{
		itemRepository: NewMemoryItemRepository(),
	}
	r.Get("/items/{id}", handler.getItemById)
	r.Get("/items", handler.getItems)
	http.ListenAndServe(":8000", r)
}
