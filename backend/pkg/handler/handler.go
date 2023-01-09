package handler

import (
	"net/http"
)

type Handler interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	// Other resource handler functions go here...
}