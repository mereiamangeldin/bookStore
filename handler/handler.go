package handler

import (
	"github.com/gorilla/mux"
	"github.com/mereiamangeldin/bookStore/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", h.createBook).Methods("POST")
	r.HandleFunc("/books", h.getAllBooks).Methods("GET")
	r.HandleFunc("/books/{id}", h.getBookById).Methods("GET")
	r.HandleFunc("/books/{id}", h.updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", h.deleteBook).Methods("DELETE")
	return r
}
