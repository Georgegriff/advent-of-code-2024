package diskmap

import "errors"

type OrderedSet[T comparable] struct {
	elements []T
	indexMap map[T]int
}

func NewOrderedSet[T comparable]() *OrderedSet[T] {
	return &OrderedSet[T]{
		elements: []T{},
		indexMap: make(map[T]int),
	}
}

func (s *OrderedSet[T]) Add(value T) {
	if _, exists := s.indexMap[value]; !exists {
		s.indexMap[value] = len(s.elements)
		s.elements = append(s.elements, value)
	}
}

func (s *OrderedSet[T]) Remove(value T) {
	index, exists := s.indexMap[value]
	if exists {
		s.elements = append(s.elements[:index], s.elements[index+1:]...)
		delete(s.indexMap, value)
		for i := index; i < len(s.elements); i++ {
			s.indexMap[s.elements[i]] = i
		}
	}
}

func (s *OrderedSet[T]) Contains(value T) bool {
	_, exists := s.indexMap[value]
	return exists
}

func (s *OrderedSet[T]) Size() int {
	return len(s.elements)
}

func (s *OrderedSet[T]) Nth(index int) (T, error) {
	if index < 0 || index >= len(s.elements) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return s.elements[index], nil
}
