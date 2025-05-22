package router

import (
	"QuotationBook/handler"
	"QuotationBook/model"
	"github.com/gorilla/mux"
)

func RoutInit(store *model.Store) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", handler.CreateQuote(store)).Methods("POST")
	r.HandleFunc("/quotes", handler.GetQuotes(store)).Methods("GET")
	r.HandleFunc("/quotes/random", handler.GetRandomQuote(store)).Methods("GET")
	r.HandleFunc("/quotes/{id}", handler.DeleteQuote(store)).Methods("DELETE")

	return r
}
