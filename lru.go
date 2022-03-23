package lru

import (
	"errors"
)

var (
	InvalidStorageSize error = errors.New("invalid storage size, must provide a positive value")
)

type LRU[T any] struct {
	capacity int
}

func New[T any](capacity int) (*LRU[T], error) {
	if capacity <= 0 {
		return nil, InvalidStorageSize
	}

	c := &LRU[T]{
		capacity: capacity,
	}

	return c, nil
}
