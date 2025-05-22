package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DeleteQuote(store *model.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Не правильный id", http.StatusBadRequest)
			return
		}
		if !store.Delete(id) {
			http.Error(w, "Не существует такого id", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := struct {
			Message string `json:"message"`
			ID      int    `json:"id"`
		}{
			Message: "Запись удалена",
			ID:      id,
		}
		json.NewEncoder(w).Encode(response)
	}
}
