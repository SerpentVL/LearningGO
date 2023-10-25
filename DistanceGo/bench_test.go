package main

import "testing"

func BenchmarkDoAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPar()
	}
}
