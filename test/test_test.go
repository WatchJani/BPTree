package test

import (
	"testing"
)

// 5.529 ns/op
// func BenchmarkInsertSlice1(b *testing.B) {
// 	b.StopTimer()
// 	data := []int{1, 2, 3, 4, 5, 7, 0}

// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		InsertSlice(data, 6, 6)
// 	}
// }

// 4.740 ns/op
func BenchmarkInsertSlice2(b *testing.B) {
	b.StopTimer()
	data := []int{1, 2, 3, 4, 5, 7, 0}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		InsertSlice2(data, 6, 6)
	}
}

// 2.767 ns/op
func BenchmarkInsertSlice3(b *testing.B) {
	b.StopTimer()
	data := []int{1, 2, 3, 4, 5, 7, 0}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		InsertSlice3(data, 6, 6)
	}
}

// 0.2158 ns/op
func BenchmarkReset1(b *testing.B) {
	b.StopTimer()
	data := make([]int, 150)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		data = data[:0]
	}
}

// 149.6 ns/op
func BenchmarkReset2(b *testing.B) {
	b.StopTimer()
	date := make([]int, 150)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		date = make([]int, 150)
	}

	b.StopTimer()
	date = date[:0]
}

// 0 alocation
func BenchmarkAllocationMemory(b *testing.B) {
	b.StopTimer()
	date := make([]int, 150)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		Memory(date)
	}
}
