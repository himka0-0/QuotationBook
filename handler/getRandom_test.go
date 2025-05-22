package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRandomQuote_Success(t *testing.T) {
	store := model.NewStore()
	added := store.Add("Author A", "Something wise")

	req := httptest.NewRequest("GET", "/quotes/random", nil)
	rec := httptest.NewRecorder()

	handler := GetRandomQuote(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	var q model.Quote
	if err := json.NewDecoder(rec.Body).Decode(&q); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if q.ID != added.ID || q.Author != added.Author {
		t.Errorf("unexpected quote: %+v", q)
	}
}

func TestGetRandomQuote_EmptyStore(t *testing.T) {
	store := model.NewStore()

	req := httptest.NewRequest("GET", "/quotes/random", nil)
	rec := httptest.NewRecorder()

	handler := GetRandomQuote(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Нет цитат") {
		t.Errorf("expected error message 'Нет цитат', got: %s", rec.Body.String())
	}
}
