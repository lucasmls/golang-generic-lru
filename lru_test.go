package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Run("Failure tests", func(t *testing.T) {
		tt := []struct {
			name     string
			capacity int
			want     error
		}{
			{
				name:     "Should not be able to instantiate a LRU with no storage capacity",
				want:     InvalidStorageSize,
				capacity: 0,
			},
		}
		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				_, err := New[string](tc.capacity)

				assert.EqualError(t, err, tc.want.Error())
			})
		}
	})

	t.Run("Successful tests", func(t *testing.T) {
		tt := []struct {
			name     string
			capacity int
			want     *LRU[string]
		}{
			{
				name: "Should be able to instantiate a LRU with the correct storage capacity",
				want: &LRU[string]{
					capacity: 10,
					storage:  make(map[string]string, 10),
				},
				capacity: 10,
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				cache, err := New[string](tc.capacity)

				assert.NoError(t, err)
				assert.Equal(t, tc.want, cache)
			})
		}
	})
}

func Test_Add(t *testing.T) {
	tt := []struct {
		name  string
		key   string
		value string
	}{
		{
			name:  "Should be able to include a new entry into cache",
			key:   "users-1",
			value: "John",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cache, err := New[string](3)
			assert.NoError(t, err)

			cache.Add(tc.key, tc.value)
			entry, ok := cache.Get(tc.key)

			assert.True(t, ok)
			assert.Equal(t, tc.value, entry)
		})
	}
}

func Test_Get(t *testing.T) {
	t.Run("Failure tests", func(t *testing.T) {
		tt := []struct {
			name string
			key  string
			want bool
		}{
			{
				name: "Shouldn't be able to retrieve a non previously added entry from cache",
				key:  "users-1",
				want: false,
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				cache, err := New[string](3)
				assert.NoError(t, err)

				_, ok := cache.Get(tc.key)
				assert.Equal(t, tc.want, ok)
			})
		}
	})

	t.Run("Successful tests", func(t *testing.T) {
		tt := []struct {
			name       string
			key        string
			want       string
			beforeEach func(*LRU[string])
		}{
			{
				name: "Should be able to retrieve a previously added entry from cache",
				key:  "users-1",
				want: "John",
				beforeEach: func(cache *LRU[string]) {
					cache.Add("users-1", "John")
				},
			},
		}

		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				cache, err := New[string](3)
				assert.NoError(t, err)

				if tc.beforeEach != nil {
					tc.beforeEach(cache)
				}

				entry, ok := cache.Get(tc.key)
				assert.True(t, ok)
				assert.Equal(t, tc.want, entry)
			})
		}
	})
}
