package leetcode

import (
	"math"
	"math/rand"
	"strconv"
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

func findWords(board [][]byte, words []string) []string {
	used := make([][]bool, len(board))
	for i := range used {
		used[i] = make([]bool, len(board[0]))
	}

	var dfs func(i, j, step int, word string) bool
	dfs = func(i, j, step int, word string) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || used[i][j] {
			return false
		}

		if step == len(word)-1 {
			if board[i][j] == word[step] {
				return true
			} else {
				return false
			}
		}

		if board[i][j] != word[step] {
			return false
		}

		used[i][j] = true
		res := dfs(i-1, j, step+1, word) || dfs(i, j+1, step+1, word) || dfs(i+1, j, step+1, word) || dfs(i, j-1, step+1, word)
		used[i][j] = false
		return res
	}

	var check func(word string) bool
	check = func(word string) bool {
		for i := range board {
			for j := range board[0] {
				if dfs(i, j, 0, word) {
					return true
				}
			}
		}
		return false
	}
	ans := make([]string, 0)
	for _, v := range words {
		if check(v) {
			ans = append(ans, v)
		}
	}
	return ans
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dict := [26]int{}
	for i := range s {
		dict[s[i]-'a']++
	}

	for i := range t {
		dict[t[i]-'a']--
	}

	for i := range dict {
		if dict[i] != 0 {
			return false
		}
	}

	return true
}

func firstUniqChar(s string) int {
	dict := [26]int{}
	for i := range dict {
		dict[i] = -1 // 未出现过
	}
	for i := range s {
		if dict[s[i]-'a'] == -1 {
			dict[s[i]-'a'] = i
		} else {
			dict[s[i]-'a'] = -2 // 出现多次
		}
	}

	ans := len(s)
	for i := range dict {
		if dict[i] != -1 && dict[i] != -2 {
			ans = min(ans, dict[i])
		}
	}

	if ans == len(s) {
		ans = -1
	}
	return ans
}

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func maxProduct(nums []int) int {
	ans, maxV, minV := nums[0], 1, 1
	for i := range nums {
		if nums[i] < 0 {
			maxV, minV = minV, maxV
		}
		maxV = max(nums[i], maxV*nums[i])
		minV = min(nums[i], minV*nums[i])
		ans = max(maxV, ans)
	}
	return ans
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}

	var reverse func(l, r int)
	reverse = func(l, r int) {
		i, j := l, r
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)
}

func containsDuplicate(nums []int) bool {
	dict := make(map[int]bool)

	for i := range nums {
		if dict[nums[i]] {
			return true
		}
		dict[nums[i]] = true
	}
	return false
}

func moveZeroes(nums []int) {
	j := -1
	for i := range nums {
		if nums[i] != 0 {
			j++
			nums[j] = nums[i]
		}
	}

	for i := j + 1; i < len(nums); i++ {
		nums[i] = 0
	}
}

func intersect(nums1 []int, nums2 []int) []int {
	dict := make(map[int]int)
	for i := range nums1 {
		dict[nums1[i]]++
	}

	ans := make([]int, 0)
	for i := range nums2 {
		if dict[nums2[i]] > 0 {
			ans = append(ans, nums2[i])
			dict[nums2[i]]--
		}
	}
	return ans
}

func findKthLargest(nums []int, k int) int {
	var partition func(l, r int) int
	partition = func(l, r int) int {
		i := rand.Intn(r-l+1) + l
		nums[i], nums[r] = nums[r], nums[i]

		lt, gt, j := l-1, r, l
		for j < gt {
			if nums[j] < nums[r] {
				nums[j], nums[lt+1] = nums[lt+1], nums[j]
				j++
				lt++
			} else if nums[j] > nums[r] {
				nums[j], nums[gt-1] = nums[gt-1], nums[j]
				gt--
			} else {
				j++
			}
		}
		nums[gt], nums[r] = nums[r], nums[gt]
		return gt
	}

	var quickSelect func(l, r, index int) int
	quickSelect = func(l, r, index int) int {
		q := partition(l, r)
		if q == len(nums)-index {
			return nums[q]
		} else if q > len(nums)-index {
			return quickSelect(l, q-1, index)
		} else {
			return quickSelect(q+1, r, index)
		}
	}

	return quickSelect(0, len(nums)-1, k)
}

func evalRPN(tokens []string) int {
	nums := make([]int, 0)

	for _, v := range tokens {
		switch v {
		case "+":
			nums[len(nums)-2] = nums[len(nums)-2] + nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		case "-":
			nums[len(nums)-2] = nums[len(nums)-2] - nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		case "*":
			nums[len(nums)-2] = nums[len(nums)-2] * nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		case "/":
			nums[len(nums)-2] = nums[len(nums)-2] / nums[len(nums)-1]
			nums = nums[:len(nums)-1]
		default:
			num, _ := strconv.Atoi(v)
			nums = append(nums, num)
		}
	}

	ans := 0
	for i := range nums {
		ans += nums[i]
	}
	return ans
}

func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)

	cur := head
	for cur != nil {
		hash[cur] = &Node{Val: cur.Val}
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		hash[cur].Next = hash[cur.Next]
		hash[cur].Random = hash[cur.Random]
		cur = cur.Next
	}

	return hash[head]
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func sortList(head *ListNode) *ListNode {
}
