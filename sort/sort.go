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
	twoWayQuickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l, l-1
	for ; i < r; i++ {
		if arr[i] < arr[r] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[j+1], arr[r] = arr[r], arr[j+1]

	quickSort(arr, l, j)
	quickSort(arr, j+2, r)
}

// 双路快排
func twoWayQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l, r-1
	for {
		for i < r && arr[i] < arr[r] {
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
	twoWayQuickSort(arr, l, i-1)
	twoWayQuickSort(arr, i+1, r)
}

// 三路快排
func threeWayQuickSort(arr []int, l, r int) {
	if l >= r {
		return
	}

	lt, i, gt := l-1, l, r
	for i < gt {
		if arr[i] < arr[r] {
			arr[i], arr[lt+1] = arr[lt+1], arr[i]
			i++
			lt++
		} else if arr[i] > arr[r] {
			arr[i], arr[gt-1] = arr[gt-1], arr[i]
			gt--
		} else {
			i++
		}
	}
	arr[r], arr[gt] = arr[gt], arr[r]
	threeWayQuickSort(arr, l, lt)
	threeWayQuickSort(arr, gt+1, r)
}

// 归并排序
func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		p := l + (r-l)>>1
		mergeSort(arr, l, p)
		mergeSort(arr, p+1, r)
		merge(arr, l, p, r)
	}
}

func merge(arr []int, l, p, r int) {
	temp := make([]int, 0, r-l+1)
	i, j, m, n := l, p+1, p, r
	for i <= m && j <= n {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else if arr[i] > arr[j] {
			temp = append(temp, arr[j])
			j++
		}
	}

	for i <= m {
		temp = append(temp, arr[i])
		i++
	}

	for j <= n {
		temp = append(temp, arr[j])
		j++
	}

	for i := range temp {
		arr[l+i] = temp[i]
	}
}

// 计数排序（不稳定版本）
func CountingSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	minVal, maxVal := arr[0], arr[0]
	for i := range arr {
		if arr[i] > maxVal {
			maxVal = arr[i]
		} else if arr[i] < minVal {
			minVal = arr[i]
		}
	}

	counts := make([]int, maxVal-minVal+1)
	for i := range arr {
		counts[arr[i]-minVal]++
	}

	sortIndex := 0
	for i := range counts {
		for counts[i] > 0 {
			arr[sortIndex] = minVal + i
			sortIndex++
			counts[i]--
		}
	}
}
