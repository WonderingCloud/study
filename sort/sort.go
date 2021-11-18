package sort

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// BubbleSortOpt1 冒泡排序第一版优化
func BubbleSortOpt1(arr []int) {
	end := len(arr) - 1
	sorted := false

	for !sorted && end > 0 {
		sorted = true
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
		end--
	}
}

// BubbleSortOpt2 冒泡排序第二版优化
func BubbleSortOpt2(arr []int) {
	end := len(arr) - 1
	
	for end > 0 {
		swapIndex := 0
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapIndex = i
			}
		}
		end = swapIndex
	}
}
