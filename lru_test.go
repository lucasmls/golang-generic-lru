package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
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
				},
				capacity: 10,
			},
		}
		for _, tc := range tt {
			t.Run(tc.name, func(t *testing.T) {
				cache, err := New[string](tc.capacity)

				assert.NoError(t, err)
				assert.Equal(t, cache, tc.want)
			})
		}
	})
}
