package handler

import (
	"encoding/json"
	"fmt"
	"github.com/mereiamangeldin/bookStore/models"
	"log"
	"net/http"
	"net/url"
)

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	res := h.service.CreateBook(book)
	if res != nil {
		http.Error(w, fmt.Sprintf("%s", res), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{Message: "Book created successfully"})
}

func (h *Handler) getAllBooks(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("title")
	if key != "" {
		key, err := url.QueryUnescape(key)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
		book, err := h.service.IAllBooks.GetBookByTitle(key)
		if err != nil {
			http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
			return
		}
		err = json.NewEncoder(w).Encode(book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	sort := r.URL.Query().Get("sort")
	books, err := h.service.IAllBooks.GetAllBooks(sort)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
