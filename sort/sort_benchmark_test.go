package sort

import (
	"testing"
)

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSortOpt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		BubbleSortOpt1(arr)
	}
}

func BenchmarkBubbleSortOpt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		BubbleSortOpt2(arr)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		QuickSort(arr)
	}
}
