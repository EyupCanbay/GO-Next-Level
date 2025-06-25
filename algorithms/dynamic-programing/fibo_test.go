package main

import "testing"

// BenchmarkFib-12    	  444591	      2639 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFib1(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = Fib1(15)
	}
}

// BenchmarkFib2-12    	829090004	         1.416 ns/op	       0 B/op	       0 allocs/op
func BenchmarkFib2(b *testing.B) {
	n := 15
	memo := make([]int, n)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fib2(n, memo)
	}
}

// BenchmarkFib3-12    	168202659	         6.929 ns/op	       0 B/op	       0 allocs/op

func BenchmarkFib3(b *testing.B) {
	n := 15
	memo := make([]int, n+1)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Fib3(n, memo)
	}
}
