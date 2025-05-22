package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
)

func GetRandomQuote(store *model.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := store.GetRandom()
		if err != nil {
			http.Error(w, "Нет цитат", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(q)
	}
}
