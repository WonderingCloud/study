package array

import "sort"

func searchInsert(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		mid := i + (j-i)>>1
		if nums[mid] >= target {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func peakIndexInMountainArray(arr []int) int {
	i, j := 0, len(arr)-1
	for i < j {
		mid := i + (j-i)>>1
		if arr[mid] > arr[mid+1] {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}

func mySqrt(x int) int {
	i, j := 0, x
	for i <= j {
		mid := i + (j-i)>>1
		if x/mid == mid {
			return mid
		} else if x/mid > mid {
			i = mid + 1
		} else if x/mid < mid {
			j = mid - 1
		}
	}
	return j
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := make([][]int, 0)
	interval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > interval[1] {
			ans = append(ans, interval)
			interval = intervals[i]
		} else {
			if intervals[i][1] > interval[1] {
				interval[1] = intervals[i][1]
			}
		}
	}
	ans = append(ans, interval)
	return ans
}

func sortColors(nums []int) {
	lt, i, gt := -1, 0, len(nums)
	for i < gt {
		if nums[i] == 0 {
			nums[i], nums[lt+1] = nums[lt+1], nums[i]
			lt++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[gt-1] = nums[gt-1], nums[i]
			gt--
		} else {
			i++
		}
	}
}

func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[len(ans)-1] = 1
	for i := 1; i <= rowIndex; i++ {
		for j := i + 1; j > 1; j-- {
			ans[len(ans)-j] = ans[len(ans)-j+1] + ans[len(ans)-j]
		}
	}
	return ans
}

func rotate(matrix [][]int) {
	i, j := 0, len(matrix)-1
	for i < j {
		matrix[i], matrix[j] = matrix[j], matrix[i]
		i++
		j--
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}

	r, b, l, t := n-1, n-1, 0, 0
	num := 1
	for num <= n*n {
		for i := l; i <= r; i++ {
			ans[t][i] = num
			num++
		}
		t++

		for i := t; i <= b; i++ {
			ans[i][r] = num
			num++
		}
		r--

		for i := r; i >= l; i-- {
			ans[b][i] = num
			num++
		}
		b--

		for i := b; i >= t; i-- {
			ans[i][l] = num
			num++
		}
		l++
	}
	return ans
}

func searchMatrix(matrix [][]int, target int) bool {
	m, n := 0, len(matrix[0])-1
	for m < len(matrix) && n >= 0 {
		if matrix[m][n] == target {
			return true
		} else if matrix[m][n] > target {
			n--
		} else if matrix[m][n] < target {
			m++
		}
	}
	return false
}

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	minVal, maxVal := nums[0], nums[0]
	for i := range nums {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}

		if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	hash := make([]int, maxVal-minVal+1)
	for i := range nums {
		hash[nums[i]-minVal]++
	}

	curIndex := 0
	for i := range hash {
		for hash[i] > 0 {
			nums[curIndex] = minVal+i
			curIndex++
		}
	}
	return nums
}
