package model

import (
	"testing"
)

func TestStore_AddAndGetAll(t *testing.T) {
	store := NewStore()

	store.Add("Author A", "Quote A")
	store.Add("Author B", "Quote B")

	all := store.GetAll("")
	if len(all) != 2 {
		t.Errorf("expected 2 quotes, got %d", len(all))
	}

	filtered := store.GetAll("Author A")
	if len(filtered) != 1 {
		t.Errorf("expected 1 quote for Author A, got %d", len(filtered))
	}
}

func TestStore_Delete(t *testing.T) {
	store := NewStore()
	q := store.Add("A", "To be deleted")

	if !store.Delete(q.ID) {
		t.Errorf("expected deletion to succeed")
	}
	if store.Delete(q.ID) {
		t.Errorf("expected second deletion to fail (already deleted)")
	}
}

func TestStore_GetRandom(t *testing.T) {
	store := NewStore()
	_, err := store.GetRandom()
	if err == nil {
		t.Errorf("expected error when store is empty")
	}

	store.Add("A", "Random quote")
	q, err := store.GetRandom()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if q.ID == 0 {
		t.Errorf("expected valid quote")
	}
}
