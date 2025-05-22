package model

import "sync"

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type Store struct {
	mu     sync.Mutex
	data   map[int]Quote
	nextID int
}

func NewStore() *Store {
	return &Store{
		data:   make(map[int]Quote),
		nextID: 1,
	}
}
