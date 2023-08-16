package cache

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Usage(t *testing.T) {
	m, _ := New(NewHashSumPartitioner(5))
	err := m.Set("testKey", "testValue")
	if assert.Nil(t, err) {
		gotValue, gotErr := m.Get("testKey")
		if assert.Nil(t, gotErr) {
			assert.Equal(t, "testValue", gotValue)
		}
	}
	_, gotErr := m.Get("xyz")
	assert.Equal(t, ErrKeyNotFound, gotErr)
}

func TestCache_New(t *testing.T) {
	DefaultCache, _ := New(NewHashSumPartitioner(DefaultNumPartitions))

	type result struct {
		cache *Cache
		err   error
	}

	type testCase struct {
		name        string
		partitioner partitioner
		expected    result
	}

	testCases := []testCase{
		{"nil partitioner", nil, result{DefaultCache, nil}},
		{"zero num partitions",
			NewHashSumPartitioner(0),
			result{nil, ErrInvalidNumPartitions},
		},
		{"one partition",
			NewHashSumPartitioner(1),
			result{
				&Cache{
					partitions: []*partition{
						{data: make(map[string]any)},
					},
					partitioner: NewHashSumPartitioner(1),
				},
				nil,
			},
		},
		{"three partitions",
			NewHashSumPartitioner(3),
			result{
				&Cache{
					partitions: []*partition{
						{data: make(map[string]any)},
						{data: make(map[string]any)},
						{data: make(map[string]any)},
					},
					partitioner: NewHashSumPartitioner(3),
				},
				nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got result
			got.cache, got.err = New(tc.partitioner)
			assert.Equal(t, tc.expected, got)
			if tc.expected.cache != nil || tc.expected.err == nil {
				assert.Equal(t, uint(cap(tc.expected.cache.partitions)), got.cache.partitioner.NumParts())
				assert.Equal(t, uint(len(tc.expected.cache.partitions)), got.cache.partitioner.NumParts())
			}
		})
	}
}
