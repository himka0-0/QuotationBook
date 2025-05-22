package handler

import (
	"QuotationBook/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateQuote(t *testing.T) {
	store := model.NewStore()
	handler := CreateQuote(store)

	body := []byte(`{"author":"Test","quote":"Testing create"}`)
	req, _ := http.NewRequest("POST", "/quotes", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201 Created, got %d", rec.Code)
	}

	var resp struct {
		Message string      `json:"message"`
		Data    model.Quote `json:"data"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if resp.Data.Author != "Test" || resp.Data.Quote != "Testing create" {
		t.Errorf("unexpected quote data: %+v", resp.Data)
	}
}
