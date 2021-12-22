package array

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
