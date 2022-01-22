package slidingwindow

import (
	"math/rand"
)

func findLengthOfLCIS(nums []int) int {
	start, ans := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ans = max(ans, i-start+1)
		} else {
			start = i
		}
	}
	return ans
}

// 80. 删除有序数组中的重复项 II
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}

	j := 1
	for i := 2; i < l; i++ {
		if nums[i] != nums[j-1] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// 75. 颜色分类
func sortColors(nums []int) {
	l, r, i := -1, len(nums), 0
	for i < r {
		if nums[i] < 1 {
			l++
			nums[i], nums[l] = nums[l], nums[i]
			i++
		} else if nums[i] > 1 {
			r--
			nums[i], nums[r] = nums[r], nums[i]
		} else {
			i++
		}
	}
}

// 215. 数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	partition := func(l, r int) int {
		idx := rand.Intn(r-l+1) + l
		nums[idx], nums[r] = nums[r], nums[idx]

		lt, rt, i := l-1, r, l
		for i < rt {
			if nums[i] < nums[r] {
				lt++
				nums[i], nums[lt] = nums[lt], nums[i]
				i++
			} else if nums[i] > nums[r] {
				rt--
				nums[i], nums[rt] = nums[rt], nums[i]
			} else {
				i++
			}
		}
		nums[rt], nums[r] = nums[r], nums[rt]
		return rt
	}

	var quickSelect func(l, r int) int
	quickSelect = func(l, r int) int {
		q := partition(l, r)
		if q == len(nums)-k {
			return nums[q]
		} else if q > len(nums)-k {
			return quickSelect(l, q-1)
		} else {
			return quickSelect(q+1, r)
		}
	}

	return quickSelect(0, len(nums)-1)
}

// 643. 子数组最大平均数 I
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum, j := sum, 0
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[j]
		maxSum = max(maxSum, sum)
		j++
	}
	return float64(maxSum) / float64(k)
}

// 1052. 爱生气的书店老板
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	j, res, extra, temp := 0, 0, 0, 0
	for i := range customers {
		if i-j == minutes {
			temp -= grumpy[j] * customers[j]
			j++
		}
		res += (1 - grumpy[i]) * customers[i]
		temp += grumpy[i] * customers[i]
		extra = max(extra, temp)
	}
	return res + extra
}

// 1423. 可获得的最大点数
func maxScore(cardPoints []int, k int) int {
	sum := 0
	for i := len(cardPoints) - k; i < len(cardPoints); i++ {
		sum += cardPoints[i]
	}

	ans := sum
	i := len(cardPoints) - k
	j := (i + k) % len(cardPoints)
	for i != 0 {
		sum += cardPoints[j] - cardPoints[i]
		ans = max(ans, sum)
		i++
		if i == len(cardPoints) {
			i = 0
		}
		j++
		if j == len(cardPoints) {
			j = 0
		}
	}
	return ans
}

// 1456. 定长子串中元音的最大数目
func maxVowels(s string, k int) int {
	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	ans, cnt, j := 0, 0, 0
	for i := range s {
		if i-j == k {
			if vowel[s[j]] {
				cnt--
			}
			j++
		}
		if vowel[s[i]] {
			cnt++
		}
		ans = max(ans, cnt)
	}
	return ans
}

// 1658. 将 x 减到 0 的最小操作数
// 等价于求和为sum(nums)-x的最长连续子序列
func minOperations(nums []int, x int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	target := sum - x
	if target < 0 {
		return -1
	}

	if target == 0 {
		return len(nums)
	}

	sum, j, maxL := 0, -1, 0
	for i := range nums {
		sum += nums[i]

		for sum > target {
			j++
			sum -= nums[j]
		}

		if sum == target {
			maxL = max(maxL, i-j)
		}
	}

	if maxL == 0 {
		return -1
	}
	return len(nums) - maxL
}

// 76. 最小覆盖子串
func minWindow(s string, t string) string {
	dict := make(map[byte]int)
	for i := range t {
		dict[t[i]]++
	}
	bNum := len(dict)

	counts := make(map[byte]int)
	cover := 0 // 已覆盖的字符数
	j, start, minL := -1, 0, len(s)+1
	for i := range s {
		if dict[s[i]] > 0 {
			if counts[s[i]] == dict[s[i]]-1 {
				cover++
			}
			counts[s[i]]++
		}

		for cover == bNum {
			if i-j < minL {
				start = j + 1
				minL = i - j
			}
			j++
			if dict[s[j]] > 0 {
				counts[s[j]]--
				if counts[s[j]] < dict[s[j]] {
					cover--
				}
			}
		}
	}

	if minL == len(s)+1 {
		return ""
	}

	return s[start : start+minL]
}

func characterReplacement(s string, k int) int {
	dict := [26]int{}

	j, maxCount, ans := -1, 0, 0
	for i := range s {
		dict[s[i]-'A']++
		maxCount = max(maxCount, dict[s[i]-'A'])

		if i-j > maxCount+k {
			j++
			dict[s[j]-'A']--
		}
		ans = max(ans, i-j)
	}
	return ans
}

// 209. 长度最小的子数组
func minSubArrayLen(target int, nums []int) int {
	j, sum, minL := -1, 0, len(nums)+1
	for i := range nums {
		sum += nums[i]

		for sum >= target {
			minL = min(minL, i-j)
			j++
			sum -= nums[j]
		}
	}

	if minL == len(nums)+1 {
		return 0
	}
	return minL
}

// 1695. 删除子数组的最大得分
func MaximumUniqueSubarray(nums []int) int {
	index := make(map[int]int)

	sum, ans, j := 0, 0, -1
	for i := range nums {
		if k, exist := index[nums[i]]; exist {
			if k > j {
				for i := j + 1; i <= k; i++ {
					sum -= nums[i]
				}
				j = k
			}
		}
		sum += nums[i]
		index[nums[i]] = i
		ans = max(ans, sum)
	}
	return ans
}

// 438. 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	dict := [26]int{}
	for i := range p {
		dict[p[i]-'a']++
	}

	counts := [26]int{}
	ans := make([]int, 0)
	j := 0
	for i := range s {
		if dict[s[i]-'a'] > 0 {
			counts[s[i]-'a']++
		}

		if i-j == len(p)-1 {
			if counts == dict {
				ans = append(ans, j)
			}

			if dict[s[j]-'a'] > 0 {
				counts[s[j]-'a']--
			}
			j++
		}
	}
	return ans
}

// 567. 字符串的排列
func checkInclusion(s1 string, s2 string) bool {
	dict := [26]int{}
	for i := range s1 {
		dict[s1[i]-'a']++
	}

	counts := [26]int{}
	j := 0
	for i := range s2 {
		if dict[s2[i]-'a'] > 0 {
			counts[s2[i]-'a']++
		}

		if i-j == len(s1)-1 {
			if counts == dict {
				return true
			}

			if dict[s2[j]-'a'] > 0 {
				counts[s2[j]-'a']--
			}
			j++
		}
	}
	return false
}
