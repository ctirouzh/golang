package cache

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkCache_BuiltinMap(b *testing.B) {
	simpleMap := make(map[string]int)
	b.Run("set", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			simpleMap[fmt.Sprint(i)] = i
		}
	})
	b.Run("set concurrently", func(b *testing.B) {
		b.ReportAllocs()
		var wg sync.WaitGroup
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(n int) {
				mu.Lock()
				simpleMap[fmt.Sprint(n)] = n
				mu.Unlock()
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = simpleMap[fmt.Sprint(i)]
		}
	})
	b.Run("get concurrently", func(b *testing.B) {
		b.ReportAllocs()
		var wg sync.WaitGroup
		var mu sync.RWMutex
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(n int) {
				mu.RLock()
				_, _ = simpleMap[fmt.Sprint(n)]
				mu.RUnlock()
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func BenchmarkCache_SyncMap(b *testing.B) {
	b.Run("set concurrently", func(b *testing.B) {
		b.ReportAllocs()
		var syncMap sync.Map
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(n int) {
				syncMap.Store(fmt.Sprint(n), n)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}

func BenchmarkCache_Partitioned(b *testing.B) {
	partitionedMap, _ := New(NewHashSumPartitioner(1000))
	b.Run("set", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = partitionedMap.Set(fmt.Sprint(i), i)
		}
	})
	b.Run("set concurrently", func(b *testing.B) {
		b.ReportAllocs()
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(n int) {
				_ = partitionedMap.Set(fmt.Sprint(n), n)
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
	b.Run("get", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, _ = partitionedMap.Get(fmt.Sprint(i))
		}
	})
	b.Run("get concurrently", func(b *testing.B) {
		b.ReportAllocs()
		var wg sync.WaitGroup
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func(n int) {
				_, _ = partitionedMap.Get(fmt.Sprint(n))
				wg.Done()
			}(i)
		}
		wg.Wait()
	})
}
