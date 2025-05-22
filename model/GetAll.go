package model

func (s *Store) GetAll(author string) []Quote {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []Quote
	for _, q := range s.data {
		if author == "" || q.Author == author {
			result = append(result, q)
		}
	}
	return result
}
