package leetcode

import (
	"math"
	"sort"
	"strconv"
)

type void struct{}

// 1. 两数之和（哈希表）
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := range nums {
		if j, exist := hash[target-nums[i]]; exist {
			return []int{j, i}
		}
		hash[nums[i]] = i
	}
	return nil
}

// 2. 两数相加（链表）
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, l1}
	l1 = dummy
	carry, sum := 0, 0
	for l1.Next != nil && l2 != nil {
		sum = l1.Next.Val + l2.Val + carry
		l1.Next.Val = sum % 10
		carry = sum / 10
		l1 = l1.Next
		l2 = l2.Next
	}

	if l1.Next == nil && l2 != nil {
		l1.Next = l2
	}

	for l1.Next != nil {
		sum = l1.Next.Val + carry
		l1.Next.Val = sum % 10
		carry = sum / 10
		l1 = l1.Next
	}

	if carry == 1 {
		l1.Next = &ListNode{1, nil}
	}
	return dummy.Next
}

// 3. 无重复字符的最长子串（滑动窗口）
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	i, j, ans := 0, -1, 0
	for ; i < len(s); i++ {
		if k, exist := hash[s[i]]; exist {
			j = max(j, k)
		}
		ans = max(ans, i-j)
		hash[s[i]] = i
	}
	return ans
}

// 4. 寻找两个正序数组的中位数（双指针、归并）
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	mid := l >> 1
	a, b, i, j, k := 0, 0, 0, 0, 0
	for k <= mid && i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			a, b = b, nums1[i]
			i++
		} else {
			a, b = b, nums2[j]
			j++
		}
		k++
	}

	for k <= mid && i < len(nums1) {
		a, b = b, nums1[i]
		i++
		k++
	}

	for k <= mid && j < len(nums2) {
		a, b = b, nums2[j]
		j++
		k++
	}

	ans := float64(0)
	if l%2 != 0 {
		ans = float64(b)
	} else {
		ans = float64(a+b) / float64(2)
	}
	return ans
}

// 5. 最长回文子串（动态规划）
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	begin, maxLen := 0, 1
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := l + i - 1
			if j >= len(s) {
				break
			}

			if s[i] == s[j] {
				if l <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] {
				begin = i
				maxLen = l
			}
		}
	}
	return s[begin : begin+maxLen]
}

// 6. Z 字形变换
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	bytes := make([][]byte, numRows)
	row, i := -1, 0
	for i < len(s) {
		for i < len(s) && row < numRows-1 {
			row++
			bytes[row] = append(bytes[row], s[i])
			i++
		}

		for i < len(s) && row > 0 {
			row--
			bytes[row] = append(bytes[row], s[i])
			i++
		}
	}
	ans := make([]byte, 0, len(s))
	for i := range bytes {
		ans = append(ans, bytes[i]...)
	}
	return string(ans)
}

// 7. 整数反转
func reverse(x int) int {
	ans := 0
	for x != 0 {
		ans = 10*ans + x%10
		x /= 10
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		ans = 0
	}
	return ans
}

// 9. 回文数
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 10. 正则表达式匹配
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}

// 11. 盛最多水的容器（贪心、双指针）
func maxArea(height []int) int {
	i, j, ans, temp := 0, len(height)-1, 0, 0
	for i < j {
		if height[i] < height[j] {
			temp = height[i] * (j - i)
			i++
		} else {
			temp = height[j] * (j - i)
			j--
		}
		ans = max(ans, temp)
	}
	return ans
}

// 12. 整数转罗马数字
func intToRoman(num int) string {
	bins := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	signs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for i := range bins {
		for num >= bins[i] {
			for j := 0; j < num/bins[i]; j++ {
				ans += signs[i]
			}
			num = num % bins[i]
		}
	}
	return ans
}

// 14. 最长公共前缀
func longestCommonPrefix(strs []string) string {
	ans := make([]byte, 0)
	s := strs[0]
	for i := range s {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != s[i] {
				return string(ans)
			}
		}
		ans = append(ans, s[i])
	}
	return string(ans)
}

// 15. 三数之和
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j+1] == nums[j] {
					j++
				}
				for j < k && nums[k-1] == nums[k] {
					k--
				}
				j++
				k--
			} else if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return ans
}

// 16. 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[2]
	sum := 0
	for i := range nums {
		j, k := i+1, len(nums)-1
		for j < k {
			sum = nums[i] + nums[j] + nums[k]
			if abs(sum-target) < abs(ans-target) {
				ans = sum
			}
			if sum == target {
				return sum
			} else if sum > target {
				k--
			} else if sum < target {
				j++
			}
		}
	}
	return ans
}

// 17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return nil
	}
	hash := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	ans := make([]string, 0)
	dfs := func(ans *[]string, s *string, step int) {}
	dfs = func(ans *[]string, s *string, step int) {
		if step == l {
			*ans = append(*ans, *s)
			return
		}

		for _, v := range hash[digits[step]] {
			*s = *s + string(v)
			dfs(ans, s, step+1)
			*s = (*s)[:step]
		}
	}
	s := ""
	dfs(&ans, &s, 0)
	return ans
}

// 18. 四数之和
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	sum := 0
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, len(nums)-1
			for k < l {
				sum = nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					for k < l && nums[k+1] == nums[k] {
						k++
					}
					for k < l && nums[l-1] == nums[l] {
						l--
					}
					k++
					l--
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return ans
}

// 19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	l, r := dummy, head
	for i := 0; i < n; i++ {
		r = r.Next
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next
	return dummy.Next
}

// 20. 有效的括号
func isValid(s string) bool {
	hash := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := make([]byte, 0)
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != hash[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

// 22. 括号生成（回溯）
func generateParenthesis(n int) []string {
	hash := make(map[string]void)
	ans := make([]string, 0)
	dfs := func(ans *[]string, s *string, step int) {}
	dfs = func(ans *[]string, s *string, step int) {
		if step == n {
			*ans = append(*ans, *s)
			return
		}

		for i := 0; i <= len(*s); i++ {
			*s = (*s)[0:i] + "()" + (*s)[i:]
			if _, exist := hash[*s]; exist {
				*s = (*s)[0:i] + (*s)[i+2:]
				continue
			}
			hash[*s] = void{}
			dfs(ans, s, step+1)
			*s = (*s)[0:i] + (*s)[i+2:]
		}
	}
	s := ""
	dfs(&ans, &s, 0)
	return ans
}

// 24. 两两交换链表中的节点
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next.Next
		cur.Next.Next = cur.Next.Next.Next
		temp.Next = cur.Next
		cur.Next = temp
		cur = cur.Next.Next
	}
	return dummy.Next
}

// 26. 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for ; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

// 27. 移除元素
func removeElement(nums []int, val int) int {
	i, j := 0, -1
	for ; i < len(nums); i++ {
		if nums[i] != val {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}

// 29. 两数相除（位运算）
func divide(dividend int, divisor int) int {
	flag := 0
	if dividend > 0 {
		dividend = -dividend
		flag++
	}
	if divisor > 0 {
		divisor = -divisor
		flag++
	}

	ans, cnt, temp := 0, 0, 0
	for dividend <= divisor {
		cnt, temp = 1, divisor
		for dividend <= temp<<1 {
			temp <<= 1
			cnt += cnt
		}
		ans += cnt
		dividend -= temp
	}

	if flag&1 == 0 && ans > math.MaxInt32 {
		ans = math.MaxInt32
	}

	if flag&1 == 1 {
		ans = -ans
	}
	return ans
}

// 704. 二分查找
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 42. 接雨水
func trap(height []int) int {
	maxHeight, maxIndex := 0, 0
	for i, h := range height {
		if h > maxHeight {
			maxHeight = h
			maxIndex = i
		}
	}

	ans, i, j, k := 0, 0, len(height)-1, 0
	for i < maxIndex {
		k = i + 1
		for height[k] < height[i] {
			ans += height[i] - height[k]
			k++
		}
		i = k
	}

	for j > maxIndex {
		k = j - 1
		for height[j] > height[k] {
			ans += height[j] - height[k]
			k--
		}
		j = k
	}
	return ans
}

// 200. 岛屿数量
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	dfs := func(i, j int) {}
	dfs = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
	}
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfs(i, j)
			}
		}
	}
	return cnt
}

// 912. 排序数组
func sortArray(nums []int) []int {
	minVal, maxVal := nums[0], nums[0]
	for i := range nums {
		if nums[i] > maxVal {
			maxVal = nums[i]
		} else if nums[i] < minVal {
			minVal = nums[i]
		}
	}

	cnts := make([]int, maxVal-minVal+1)
	for i := range nums {
		cnts[nums[i]-minVal]++
	}

	index := 0
	for i := range cnts {
		for cnts[i] > 0 {
			nums[index] = minVal + i
			index++
			cnts[i]--
		}
	}
	return nums
}

// 46. 全排列
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	used := make(map[int]struct{})
	dfs := func(step int, arr *[]int) {}
	dfs = func(step int, arr *[]int) {
		if step == len(nums) {
			copy := append([]int{}, *arr...)
			ans = append(ans, copy)
			return
		}

		for i := range nums {
			if _, exist := used[nums[i]]; !exist {
				*arr = append(*arr, nums[i])
				used[nums[i]] = struct{}{}
				dfs(step+1, arr)
				*arr = (*arr)[:step]
				delete(used, nums[i])
			}
		}
	}

	arr := make([]int, 0, len(nums))
	dfs(0, &arr)
	return ans
}

// 53. 最大子数组和
/*func maxSubArray(nums []int) int {
	ans, sum := 0, 0
	for i := range nums {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		ans = max(ans, sum)
	}
	return ans
}*/
func maxSubArray(nums []int) int {
	var maxSum func(l, r int) int
	maxSum = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		mid := l + (r-l)>>1
		leftSum := maxSum(l, mid)
		rightSum := maxSum(mid+1, r)

		val, lVal := 0, nums[mid]
		for i := mid; i >= l; i-- {
			val += nums[i]
			lVal = max(lVal, val)
		}

		val, rVal := 0, nums[mid+1]
		for i := mid + 1; i <= r; i++ {
			val += nums[i]
			rVal = max(rVal, val)
		}
		return max(lVal+rVal, max(leftSum, rightSum))
	}
	return maxSum(0, len(nums)-1)
}

// 102. 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	ans := make([][]int, 0)
	queue := []*TreeNode{root}
	l := 0
	for len(queue) > 0 {
		l = len(queue)
		level := make([]int, 0, l)
		for i := 0; i < l; i++ {
			level = append(level, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, level)
		queue = queue[l:]
	}
	return ans
}

// 47. 全排列 II
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	dfs := func(arr *[]int, used []int8, step int) {}
	dfs = func(arr *[]int, used []int8, step int) {
		if step == len(nums) {
			copy := append([]int{}, *arr...)
			ans = append(ans, copy)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] == 1 {
				continue
			}

			if i > 0 && nums[i] == nums[i-1] && used[i-1] == 0 {
				continue
			}
			used[i] = 1
			*arr = append(*arr, nums[i])
			dfs(arr, used, step+1)
			used[i] = 0
			*arr = (*arr)[:step]
		}
	}
	arr := make([]int, 0, len(nums))
	used := make([]int8, len(nums))
	dfs(&arr, used, 0)
	return ans
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	arr := make([]int, 0)
	dfs := func(target, idx int) {}
	dfs = func(target, idx int) {
		if target == 0 {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		if idx == len(candidates) {
			return
		}

		dfs(target, idx+1)
		if target-candidates[idx] >= 0 {
			arr = append(arr, candidates[idx])
			dfs(target-candidates[idx], idx)
			arr = arr[:len(arr)-1]
		}
	}
	dfs(target, 0)
	return ans
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	return nil
}

func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}

	left, right := 1, 1

	for i := range nums {
		ans[i] *= left
		ans[len(nums)-1-i] *= right
		left *= nums[i]
		right *= nums[len(nums)-1-i]
	}
	return ans
}

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	length := m * n
	ans := make([]int, length)

	l, r, t, b, i := 0, n-1, 0, m-1, 0
	for i < length {
		for j := l; i < length && j <= r; j++ {
			ans[i] = matrix[t][j]
			i++
		}
		t++

		for j := t; i < length && j <= b; j++ {
			ans[i] = matrix[j][r]
			i++
		}
		r--

		for j := r; i < length && j >= l; j-- {
			ans[i] = matrix[b][j]
			i++
		}
		b--

		for j := b; i < length && j >= t; j-- {
			ans[i] = matrix[j][l]
			i++
		}
		l++
	}
	return ans
}

func gameOfLife(board [][]int) {
	// 2: 0->1 3: 1->0
	direct := []int{-1, 0, 1}
	m, n := len(board), len(board[0])

	change := func(r, c int) {
		cnt := 0
		for _, i := range direct {
			if r+i < 0 || r+i > m-1 {
				continue
			}
			for _, j := range direct {
				if c+j < 0 || c+j > n-1 {
					continue
				}
				if i == 0 && j == 0 {
					continue
				}
				if board[r+i][c+j]&1 == 1 {
					cnt++
				}
			}
		}

		if board[r][c] == 0 {
			if cnt == 3 {
				board[r][c] = 2
			}
		} else {
			if cnt == 2 || cnt == 3 {
				board[r][c] = 1
			} else {
				board[r][c] = 3
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			change(i, j)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 1
			} else if board[i][j] == 3 {
				board[i][j] = 0
			}
		}
	}
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] >= 1 && nums[i] <= len(nums) && nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func calculate(s string) int {

	stack := make([]int, 0)

	num, op := 0, byte('+')

	operator := func(op byte, num int) {
		switch op {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -num)
		case '*':
			stack[len(stack)-1] *= num
		case '/':
			stack[len(stack)-1] /= num
		}
	}

	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
			if i == len(s)-1 {
				operator(op, num)
			}
		} else {
			if i > 0 && s[i-1] >= '0' && s[i-1] <= '9' {
				operator(op, num)
				num = 0
			}

			switch s[i] {
			case '+', '-', '*', '/':
				op = s[i]
			default:
				break
			}
		}
	}

	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum
}

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0)

	ans := make([]int, 0)
	for i := range nums {
		if i > 0 && i-queue[0] >= k {
			queue = queue[1:]
		}

		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			ans = append(ans, nums[queue[0]])
		}
	}
	return ans
}

func mergeKLists(lists []*ListNode) *ListNode {
	var mergeLists func(l, r int) *ListNode
	mergeLists = func(l, r int) *ListNode {
		if l == r {
			return lists[l]
		}

		if l > r {
			return nil
		}

		mid := l + (r-l)>>1
		return mergeTwoLists(mergeLists(l, mid), mergeLists(mid+1, r))
	}
	return mergeLists(0, len(lists)-1)
}

func solve(board [][]byte) {
	m, n := len(board), len(board[0])

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i > m-1 || j < 0 || j > n-1 {
			return
		}

		if board[i][j] == 'X' || board[i][j] == '#' {
			return
		}

		board[i][j] = '#'
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i+1, j)
		dfs(i, j-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minV, ans := prices[0], 0

	for i := 1; i < len(prices); i++ {
		ans = max(ans, max(0, prices[i]-minV))
		minV = min(minV, prices[i])
	}
	return ans
}

// 142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}

	return nil
}

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	used := make([][]bool, m)
	for i := range used {
		used[i] = make([]bool, n)
	}

	var dfs func(i, j, l int) bool
	dfs = func(i, j, l int) bool {
		if l == len(word) {
			return true
		}

		if i < 0 || i > m-1 || j < 0 || j > n-1 || used[i][j] {
			return false
		}

		if board[i][j] != word[l] {
			return false
		}

		used[i][j] = true
		res := dfs(i-1, j, l+1) || dfs(i, j+1, l+1) || dfs(i+1, j, l+1) || dfs(i, j-1, l+1)
		used[i][j] = false
		return res
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 90. 子集 II
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	used := make([]bool, len(nums))
	ans := make([][]int, 0)
	arr := make([]int, 0)
	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		dfs(idx + 1)

		if idx > 0 && nums[idx] == nums[idx-1] && !used[idx-1] {
			return
		}

		used[idx] = true
		arr = append(arr, nums[idx])
		dfs(idx + 1)
		used[idx] = false
		arr = arr[:len(arr)-1]
	}
	dfs(0)
	return ans
}

