package main

import (
	"QuotationBook/model"
	"QuotationBook/router"
	"log"
	"net/http"
)

func main() {
	store := model.NewStore()
	r := router.RoutInit(store)
	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
