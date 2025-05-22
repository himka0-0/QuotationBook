package model

import (
	"errors"
	"math/rand"
)

func (s *Store) GetRandom() (Quote, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		return Quote{}, errors.New("Нет цитат")
	}
	list := make([]Quote, 0, len(s.data))
	for _, q := range s.data {
		list = append(list, q)
	}
	return list[rand.Intn(len(list))], nil
}
