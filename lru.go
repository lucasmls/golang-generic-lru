package lru

import (
	"container/list"
	"errors"
)

var (
	InvalidStorageSize error = errors.New("invalid storage size, must provide a positive value")
)

type entry[T any] struct {
	Key   string
	Value T
}

type LRU[T any] struct {
	capacity     int
	storage      map[string]*list.Element
	evictionList list.List
}

func New[T any](capacity int) (*LRU[T], error) {
	if capacity <= 0 {
		return nil, InvalidStorageSize
	}

	var v list.List
	c := &LRU[T]{
		capacity:     capacity,
		storage:      make(map[string]*list.Element, capacity),
		evictionList: v,
	}

	return c, nil
}

func (c *LRU[T]) Add(key string, value T) (evicted bool) {
	if el, ok := c.storage[key]; ok {
		en := el.Value.(*entry[T])
		en.Value = value

		c.evictionList.MoveToFront(el)

		return false
	}

	evict := c.evictionList.Len() >= c.capacity
	if evict {
		el := c.evictionList.Back()
		if el != nil {
			c.evictionList.Remove(el)
		}

		en, ok := el.Value.(*entry[T])
		if !ok {
			return false
		}

		delete(c.storage, en.Key)
	}

	en := &entry[T]{key, value}
	el := c.evictionList.PushFront(en)
	c.storage[key] = el

	return evict
}

func (c *LRU[T]) Get(key string) (t T, ok bool) {
	el, ok := c.storage[key]
	if !ok {
		return t, false
	}

	en, ok := el.Value.(*entry[T])
	if !ok {
		return t, false
	}

	c.evictionList.MoveToFront(el)

	return en.Value, true
}
