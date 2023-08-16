package cache

import "errors"

var (
	ErrKeyNotFound          = errors.New("key not found")
	ErrInvalidNumPartitions = errors.New("number of partitions should be positive")
)
