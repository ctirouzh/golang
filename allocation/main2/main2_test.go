package main2

import "testing"

func BenchmarkMain2_StackIt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main2()
	}
}

/*
goos: linux
goarch: amd64
pkg: github.com/foldera/golang/allocation/main2
cpu: Intel(R) Core(TM) i7-6700HQ CPU @ 2.60GHz
BenchmarkMain1_StackIt
BenchmarkMain1_StackIt-8   	56508612	        19.19 ns/op	       8 B/op	       1 allocs/op
*/

/*
So does this mean pointers are guaranteed to create allocations?
Let’s modify the program again to this time pass a pointer down the stack in main3 folder...
*/
