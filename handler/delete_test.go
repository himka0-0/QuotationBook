package handler

import (
	"QuotationBook/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestDeleteQuote_Success(t *testing.T) {
	store := model.NewStore()
	q := store.Add("Test", "To be deleted")

	req := httptest.NewRequest("DELETE", "/quotes/"+strconv.Itoa(q.ID), nil)
	req = mux.SetURLVars(req, map[string]string{"id": strconv.Itoa(q.ID)})
	rec := httptest.NewRecorder()

	handler := DeleteQuote(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rec.Code)
	}

	var resp struct {
		Message string `json:"message"`
		ID      int    `json:"id"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatal("failed to decode response:", err)
	}

	if resp.ID != q.ID || resp.Message != "Запись удалена" {
		t.Errorf("unexpected response: %+v", resp)
	}
}

func TestDeleteQuote_NotFound(t *testing.T) {
	store := model.NewStore()

	req := httptest.NewRequest("DELETE", "/quotes/999", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "999"})
	rec := httptest.NewRecorder()

	handler := DeleteQuote(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 Not Found, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Не существует такого id") {
		t.Errorf("unexpected response body: %s", rec.Body.String())
	}
}

func TestDeleteQuote_InvalidID(t *testing.T) {
	store := model.NewStore()

	req := httptest.NewRequest("DELETE", "/quotes/abc", nil)
	req = mux.SetURLVars(req, map[string]string{"id": "abc"})
	rec := httptest.NewRecorder()

	handler := DeleteQuote(store)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", rec.Code)
	}

	if !strings.Contains(rec.Body.String(), "Не правильный id") {
		t.Errorf("unexpected response body: %s", rec.Body.String())
	}
}
