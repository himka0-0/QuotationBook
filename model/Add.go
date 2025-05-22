package model

func (s *Store) Add(author, text string) Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	q := Quote{
		ID:     s.nextID,
		Author: author,
		Quote:  text,
	}
	s.data[s.nextID] = q
	s.nextID++
	return q
}
