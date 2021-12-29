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

func sortColors(nums []int)  {
    lt, i, gt := -1, 0, len(nums)
    for i < gt {
        if nums[i] == 0 {
            nums[i], nums[lt+1] = nums[lt+1], nums[i]
            lt++
            i++
        } else if nums[i] == 2 {
            nums[i], nums[gt-1]= nums[gt-1], nums[i]
            gt--
        } else {
            i++
        }
    }
}
