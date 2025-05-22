package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
)

func GetQuotes(store *model.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		list := store.GetAll(author)
		w.Header().Set("Content-Type", "application/json")
		if len(list) == 0 {
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Пока нет цитат",
			})
			return
		}
		json.NewEncoder(w).Encode(list)
	}
}
