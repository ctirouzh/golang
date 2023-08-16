# **Partitioned Map (Faster Cache)**

## **Usage**

```go
package cache

func TestCache_Usage(t *testing.T) {
	m,_ := New(NewHashSumPartitioner(5))
	err := m.Set("testKey", "testValue")
	if assert.Nil(t, err) {
		gotValue, gotErr := m.Get("testKey")
		if assert.Nil(t, gotErr) {
			assert.Equal(t, "testValue", gotValue)
		}
	}
	_, gotErr := m.Get("xyz")
	assert.Equal(t, errors.New("key not found"), gotErr)
}
```

```shell
go test ./cache
```

## **Benchmarks**
```shell
go test ./cache -bench=. -benchtime=3s
```

```dtd
goos: linux
goarch: amd64
pkg: github.com/foldera/golang/cache
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz

BenchmarkCache_BuiltinMap/set-8                  6682093     661.1 ns/op    142 B/op    2 allocs/op
BenchmarkCache_Partitioned/set-8                 5466843     633.4 ns/op    117 B/op    4 allocs/op

BenchmarkCache_BuiltinMap/set_concurrently-8     3837920    1067   ns/op    147 B/op    4 allocs/op
BenchmarkCache_SyncMap/set_concurrently-8        2370961    2171   ns/op    263 B/op    9 allocs/op
BenchmarkCache_Partitioned/set_concurrently-8    8613118     489.7 ns/op    151 B/op    6 allocs/op


BenchmarkCache_BuiltinMap/get-8                 12825241     295.3 ns/op     15 B/op    1 allocs/op
BenchmarkCache_Partitioned/get-8                 9366567     418.2 ns/op     25 B/op    3 allocs/op

BenchmarkCache_BuiltinMap/get_concurrently-8     8611452     420.5 ns/op     71 B/op    3 allocs/op
BenchmarkCache_Partitioned/get_concurrently-8    8754184     441.3 ns/op     72 B/op    5 allocs/op
```



## **References**

1. [Writing a Partitioned Cache Using Go Map](https://blog.stackademic.com/writing-a-partitioned-cache-using-go-map-x3-faster-than-the-standard-map-dbfe704fe4bf)
2. [github.com/vadimInshakov/partmap](https://github.com/vadimInshakov/partmap)
