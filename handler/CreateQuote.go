package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
)

func CreateQuote(store *model.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Author string `json:"author"`
			Quote  string `json:"quote"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q := store.Add(input.Author, input.Quote)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		response := struct {
			Message string      `json:"message"`
			Data    interface{} `json:"data"`
		}{
			Message: "Запись создана",
			Data:    q,
		}
		json.NewEncoder(w).Encode(response)
	}
}
