package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mereiamangeldin/bookStore/models"
	"net/http"
	"strconv"
)

func (h *Handler) getBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	book, err := h.service.IBook.GetBookById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	res := h.service.UpdateBook(id, book)
	if res != nil {
		http.Error(w, fmt.Sprintf("%s", res), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{Message: "Book updated successfully"})
}

func (h *Handler) deleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	err = h.service.IBook.DeleteBook(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(struct{ Message string }{Message: "Book deleted successfully"})

}
