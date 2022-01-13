package leetcode

import (
	"math"
	"strings"
)

func singleNumber(nums []int) int {
	a := 0
	for _, v := range nums {
		a = a ^ v
	}
	return a
}

func majorityElement(nums []int) int {
	cnt, ans := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == ans {
			cnt++
		} else {
			if cnt > 0 {
				cnt--
			} else {
				cnt = 1
				ans = nums[i]
			}
		}
	}
	return ans
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i < 0 || nums1[i] <= nums2[j] {
			nums1[k] = nums2[j]
			j--
		} else {
			nums1[k] = nums1[i]
			i--
		}
		k--
	}
}

func superEggDrop(k int, n int) int {
	memory := make(map[[2]int]int)

	var dp func(k int, n int) int
	dp = func(k, n int) int {
		if k == 1 {
			return n
		}
		if n == 0 {
			return 0
		}

		if v, exist := memory[[2]int{k, n}]; exist {
			return v
		}

		res := math.MaxInt64
		lo, hi := 1, n
		for lo <= hi {
			mid := lo + (hi-lo)>>1
			broken := dp(k-1, mid-1)
			notBroken := dp(k, n-mid)
			if broken > notBroken {
				hi = mid - 1
				res = min(res, broken+1)
			} else {
				lo = mid + 1
				res = min(res, notBroken+1)
			}
		}
		memory[[2]int{k, n}] = res
		return res
	}
	return dp(k, n)
}

func partition(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}

			if l <= 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
		}
	}

	ans := make([][]string, 0)
	var dfs func(arr *[]string, idx int)
	dfs = func(arr *[]string, idx int) {
		if idx == len(s) {
			ans = append(ans, append([]string{}, *arr...))
			return
		}

		for i := range dp[idx] {
			if !dp[idx][i] {
				continue
			}
			*arr = append(*arr, s[idx:i+1])
			dfs(arr, i+1)
			*arr = (*arr)[:len(*arr)-1]
		}
	}
	arr := make([]string, 0)
	dfs(&arr, 0)
	return ans
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func wordBreak2(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, v := range wordDict {
		dict[v] = true
	}

	ans := make([]string, 0)
	arr := make([]string, 0)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(s) {
			ans = append(ans, strings.Join(arr, " "))
			return
		}

		for i := idx + 1; i <= len(s); i++ {
			if dict[s[idx:i]] {
				arr = append(arr, s[idx:i])
				dfs(i)
				arr = arr[:len(arr)-1]
			}
		}
	}
	dfs(0)
	return ans
}
