package main

import "testing"

const n = 10

// eg. go test ./go/benchmark -bench=Slice

func BenchmarkSliceCap(b *testing.B) {
	sl := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sl = append(sl, i)
	}
}

func BenchmarkSliceLen(b *testing.B) {
	sl := make([]int, n)
	for i := range sl {
		sl[i] = i
	}
}

func BenchmarkSliceNone(b *testing.B) {
	var sl []int
	for i := 0; i < n; i++ {
		sl = append(sl, i)
	}
}
