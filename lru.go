package lru

import (
	"errors"
)

var (
	InvalidStorageSize error = errors.New("invalid storage size, must provide a positive value")
)

type LRU[T any] struct {
	capacity int
	storage  map[string]T
}

func New[T any](capacity int) (*LRU[T], error) {
	if capacity <= 0 {
		return nil, InvalidStorageSize
	}

	c := &LRU[T]{
		capacity: capacity,
		storage:  make(map[string]T, capacity),
	}

	return c, nil
}

func (c *LRU[T]) Add(key string, value T) {
	c.storage[key] = value
}

func (c *LRU[T]) Get(key string) (T, bool) {
	value, ok := c.storage[key]
	return value, ok
}
