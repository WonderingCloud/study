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
		v, j := arr[i], i-1
		for j > -1 && arr[j] > v {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = v
	}
}

// QuickSort 快速排序
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l < r {
		p := partitionOpt1(arr, l, r)
		quickSort(arr, l, p-1)
		quickSort(arr, p+1, r)
	}
}

func partition(arr []int, l, r int) int {
	i, j := l, l-1
	for ; i < r; i++ {
		if arr[i] < arr[r] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j+1], arr[r] = arr[r], arr[j+1]
	return j + 1
}

func partitionOpt1(arr []int, l, r int) int {
	i, j := l, r-1
	for {
		for i <= r-1 && arr[i] < arr[r] {
			i++
		}

		for j >= l && arr[j] > arr[r] {
			j--
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}


func quickSortOpt(arr []int, l, r int) {

}