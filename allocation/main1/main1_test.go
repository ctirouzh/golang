package main1

import "testing"

func BenchmarkMain1_StackIt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main1()
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/foldera/golang/allocation/main1
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkMain1_StackIt
BenchmarkMain1_StackIt-8   	671195216	         1.649 ns/op	       0 B/op	       0 allocs/op
*/

/*
As expected, the allocs/op stat is 0. An important observation we can make from this result is that
copying variables can allow us to keep them on the stack and avoid allocation to the heap.
*/
