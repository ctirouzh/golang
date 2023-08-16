package cache

const (
	DefaultNumPartitions = 9
)

type Cache struct {
	partitions  []*partition
	partitioner partitioner
}

func New(partitioner partitioner) (*Cache, error) {
	if partitioner == nil {
		return New(NewHashSumPartitioner(DefaultNumPartitions))
	}
	if partitioner.NumParts() == 0 {
		return nil, ErrInvalidNumPartitions
	}
	partitions := make([]*partition, 0, partitioner.NumParts())
	for i := 0; i < int(partitioner.NumParts()); i++ {
		m := make(map[string]any)
		partitions = append(partitions, &partition{data: m})
	}

	return &Cache{partitions: partitions, partitioner: partitioner}, nil
}

func (m *Cache) Set(key string, value any) error {
	partIndex, err := m.partitioner.Find(key)
	if err != nil {
		return err
	}

	part := m.partitions[partIndex]
	part.set(key, value)

	return nil
}

func (m *Cache) Get(key string) (any, error) {
	partIndex, err := m.partitioner.Find(key)
	if err != nil {
		return nil, err
	}

	part := m.partitions[partIndex]
	value, exist := part.get(key)
	if !exist {
		return nil, ErrKeyNotFound
	}

	return value, nil
}
