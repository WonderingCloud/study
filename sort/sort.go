package mysort

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
	sorted := false
	swapIndex := 0

	for !sorted && end > 0 {
		sorted = true
		for i := 0; i < end; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapIndex = i
				sorted = false
			}
		}
		end = swapIndex
	}
}

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {

	}
}
