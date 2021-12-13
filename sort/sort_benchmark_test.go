package mysort

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 2, 3, 7, 5, 4, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSortOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 2, 3, 7, 5, 4, 8, 9, 10, 12, 11, 13, 14, 15, 16, 17}
		BubbleSortOpt1(arr)
	}
}

func BenchmarkBubbleSortOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 2, 3, 7, 5, 4, 8, 9, 10, 12, 11, 13, 14, 15, 16, 17}
		BubbleSortOpt2(arr)
	}
}
