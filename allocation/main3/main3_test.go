package main3

import "testing"

func BenchmarkMain3_StackIt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		main3()
	}
}
