package handler

import (
	"bookstore/pkg/model"
	"bookstore/pkg/repository"
	"context"
	"encoding/json"
	"net/http"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(conn *db.Conn) *BookHandler {
	return &BookHandler{repo: repository.NewBookRepository(conn)}
}

func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}

	book, err := h.repo.GetBookByID(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}