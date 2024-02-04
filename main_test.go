package main

import (
	"math/rand"
	"testing"
)

func Test(t *testing.T) {
	data := []struct {
		input   []*Key
		pointer int
		key     int
		output  int
	}{
		{[]*Key{NewKey(1), NewKey(2), NewKey(3), NewKey(4), NewKey(5), NewKey(0), NewKey(0)}, 5, 7, 5},
	}

	for index, result := range data {
		if actual := Find(result.input, result.key, result.pointer); actual != result.output {
			t.Errorf("index: %d | actual: %d | output: %d", index, actual, result.output)
		}
	}
}

// func TestInsertKey(t *testing.T) {
// 	data := []struct {
// 		input         []int
// 		pointer       int
// 		position      int
// 		key           []int
// 		output        []int
// 		pointerOutput int
// 	}{
// 		{[]int{1, 2, 3, 0, 0, 0}, 3, 3, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, 6},
// 	}

// 	for index, result := range data {
// 		if actual := Find(result.input, result.key, result.pointer); actual != result.output {
// 			t.Errorf("index: %d | actual: %d | output: %d", index, actual, result.output)
// 		}
// 	}
// }

func BenchmarkTree(b *testing.B) {
	b.StopTimer()
	tree := NewBPTree(10000, 5)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tree.Insert(rand.Intn(5000000))
	}
}
