package lru

import (
	"errors"
)

var (
	InvalidStorageSize error = errors.New("invalid storage size, must provide a positive value")
)

type entry[T any] struct {
	Value T
}

type LRU[T any] struct {
	capacity int
	storage  map[string]*entry[T]
}

func New[T any](capacity int) (*LRU[T], error) {
	if capacity <= 0 {
		return nil, InvalidStorageSize
	}

	c := &LRU[T]{
		capacity: capacity,
		storage:  make(map[string]*entry[T], capacity),
	}

	return c, nil
}

func (c *LRU[T]) Add(key string, value T) {
	c.storage[key] = &entry[T]{
		Value: value,
	}
}

func (c *LRU[T]) Get(key string) (T, bool) {
	entry, ok := c.storage[key]
	if !ok {
		var t T
		return t, false
	}

	return entry.Value, true
}
