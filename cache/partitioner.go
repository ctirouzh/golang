package cache

import "hash/crc32"

type partitioner interface {
	NumParts() uint
	Find(key string) (uint, error)
}

type HashSumPartitioner struct {
	numPartitions uint
}

func NewHashSumPartitioner(numPartitions uint) *HashSumPartitioner {
	return &HashSumPartitioner{numPartitions: numPartitions}
}

func (h *HashSumPartitioner) NumParts() uint {
	return h.numPartitions
}

func (h *HashSumPartitioner) Find(key string) (uint, error) {
	hashSum := crc32.ChecksumIEEE([]byte(key))
	return uint(hashSum) % h.numPartitions, nil
}
