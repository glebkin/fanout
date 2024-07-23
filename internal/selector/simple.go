package selector

// Simple selector acts like a queue and picks elements one-by-one starting from the first element
type Simple[T any] struct {
	values []T
	idx    int
}

func NewSimpleSelector[T any](values []T) *Simple[T] {
	return &Simple[T]{
		values: values,
		idx:    0,
	}
}

// Pick returns next available element from values array if exists.
// Returns default value of type T otherwise
func (s *Simple[T]) Pick() T {
	var result T
	if s.idx >= len(s.values) {
		return result
	}

	result = s.values[s.idx]
	s.idx++

	return result
}
