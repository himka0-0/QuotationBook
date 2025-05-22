package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestGetQuotes_AllQuotes(t *testing.T) {
	store := model.NewStore()
	store.Add("Author A", "Quote A")
	store.Add("Author B", "Quote B")

	req := httptest.NewRequest("GET", "/quotes", nil)
	rec := httptest.NewRecorder()

	handler := GetQuotes(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	var quotes []model.Quote
	if err := json.NewDecoder(rec.Body).Decode(&quotes); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if len(quotes) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(quotes))
	}
}

func TestGetQuotes_FilterByAuthor(t *testing.T) {
	store := model.NewStore()
	store.Add("Author A", "Quote A")
	store.Add("Author B", "Quote B")

	url := "/quotes?" + url.Values{"author": {"Author A"}}.Encode()
	req := httptest.NewRequest("GET", url, nil)
	rec := httptest.NewRecorder()

	handler := GetQuotes(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	var quotes []model.Quote
	if err := json.NewDecoder(rec.Body).Decode(&quotes); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if len(quotes) != 1 || quotes[0].Author != "Author A" {
		t.Errorf("filter by author failed, got: %+v", quotes)
	}
}

func TestGetQuotes_EmptyStore(t *testing.T) {
	store := model.NewStore()

	req := httptest.NewRequest("GET", "/quotes", nil)
	rec := httptest.NewRecorder()

	handler := GetQuotes(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	var resp map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	msg, ok := resp["message"]
	if !ok || msg != "Пока нет цитат" {
		t.Errorf("expected message 'Пока нет цитат', got: %v", resp)
	}
}

func TestGetQuotes_FilterNoResults(t *testing.T) {
	store := model.NewStore()
	store.Add("Author A", "Quote A")

	req := httptest.NewRequest("GET", "/quotes?author=Unknown", nil)
	rec := httptest.NewRecorder()

	handler := GetQuotes(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Пока нет цитат") {
		t.Errorf("expected message for empty result, got: %s", rec.Body.String())
	}
}
