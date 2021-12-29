package sword

import (
	"math"
	"sort"
	"strconv"
)

// 001
func divide(a int, b int) int {
	divided, temp, flag := a, b, 0
	if a > 0 {
		divided = -divided
		flag++
	}

	if b > 0 {
		temp = -temp
		flag++
	}

	ans, divisor := 0, temp
	for divided <= divisor {
		cnt := 1
		for divided <= divisor<<1 {
			divisor = divisor << 1
			cnt += cnt
		}
		ans += cnt
		divided -= divisor
		divisor = temp
	}

	if flag == 1 && ans != 0 {
		ans = -ans
	} else if flag != 1 && ans > math.MaxInt32 {
		ans = math.MaxInt32
	}

	return ans
}

// 002
func addBinary(a string, b string) string {
	x, y := []byte(a), []byte(b)
	if len(x) < len(y) {
		x, y = y, x
	}

	carry := 0
	for i := len(y) - 1; i > -1; i-- {
		sum := int(x[i+len(x)-len(y)]-'0'+y[i]-'0') + carry
		x[i+len(x)-len(y)] = byte(sum%2 + int('0'))
		carry = sum / 2
	}

	for i := len(x) - len(y) - 1; i > -1; i-- {
		sum := int(x[i]-'0') + carry
		x[i] = byte(sum%2 + int('0'))
		carry = sum / 2
	}

	if carry == 1 {
		x = append([]byte{'1'}, x...)
	}
	return string(x)
}

// 003
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		ans[i] = ans[i>>1] + i&1
	}
	return ans
}

// 004
func singleNumber(nums []int) int {
	ans := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, num := range nums {
			total += int32(num) >> i & 1
		}
		if total%3 > 0 {
			ans |= 1 << i
		}
	}
	return int(ans)
}

// 005
func maxProduct(words []string) int {
	nums := make([]int, len(words))

	for i := range words {
		for j := 0; j < len(words[i]); j++ {
			nums[i] |= 1 << int(words[i][j]-'a')
		}
	}

	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]&nums[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}

	return ans
}

// 006
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sum := numbers[i] + numbers[j]
		if sum == target {
			return []int{i, j}
		} else if sum < target {
			i++
		} else if sum > target {
			j--
		}
	}

	return nil
}

// 007
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if nums[i]+nums[left]+nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			}
		}
	}
	return res
}

// 008
func minSubArrayLen(target int, nums []int) int {
	j, sum, ans := 0, 0, len(nums)+1
	for i := range nums {
		sum += nums[i]
		for sum >= target {
			ans = min(ans, i-j+1)
			if ans == 1 {
				return ans
			}
			sum -= nums[j]
			j++
		}
	}
	if ans > len(nums) {
		ans = 0
	}
	return ans
}

// 009
func numSubarrayProductLessThanK(nums []int, k int) int {
	j, prod, ans := 0, 1, 0
	for i := range nums {
		prod *= nums[i]
		for j <= i && prod >= k {
			prod /= nums[j]
			j++
		}
		ans += i - j + 1
	}
	return ans
}

// 010
func subarraySum(nums []int, k int) int {
	hash := map[int]int{0: 1}
	sum, ans := 0, 0
	for i := range nums {
		sum += nums[i]
		ans += hash[sum-k]
		hash[sum] = hash[sum] + 1
	}
	return ans
}

// 011
func findMaxLength(nums []int) int {
	hash := map[int]int{0: -1}
	sum, ans := 0, 0
	for i := range nums {
		if nums[i] == 1 {
			sum++
		} else {
			sum--
		}
		j, ok := hash[sum]
		if ok {
			ans = max(ans, i-j)
		} else {
			hash[sum] = i
		}
	}
	return ans
}

// 012
func pivotIndex(nums []int) int {
	total := 0
	for i := range nums {
		total += nums[i]
	}

	sum := 0
	for i := range nums {
		if total == sum*2+nums[i] {
			return i
		}
		sum += nums[i]
	}
	return -1
}

// 014
func checkInclusion(s1 string, s2 string) bool {
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s1 {
		cnt1[int(s1[i]-'a')]++
	}

	j := -1
	for i := range s2 {
		cnt2[int(s2[i]-'a')]++
		if i-j == len(s1) {
			if cnt1 == cnt2 {
				return true
			}
			j++
			cnt2[int(s2[j]-'a')]--
		}
	}
	return false
}

// 015
func findAnagrams(s string, p string) []int {
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range p {
		cnt1[int(p[i]-'a')]++
	}

	j := -1
	ans := make([]int, 0)
	for i := range s {
		cnt2[int(s[i]-'a')]++
		if i-j == len(p) {
			if cnt1 == cnt2 {
				ans = append(ans, j+1)
			}
			j++
			cnt2[int(s[j]-'a')]--
		}
	}
	return ans
}

// 016
func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]int)
	j, ans := -1, 0
	for i := range s {
		if v, ok := hash[s[i]]; ok {
			j = max(j, v)
		}
		ans = max(ans, i-j)
		hash[s[i]] = i
	}
	return ans
}

// 018
func isPalindrome(s string) bool {
	bs := []byte(s)
	i, j := 0, len(bs)-1
	for i < j {
		if isalnum(bs, i) && isalnum(bs, j) {
			if bs[i] != bs[j] {
				return false
			}
			i++
			j--
		}

		if !isalnum(bs, i) {
			i++
		}

		if !isalnum(bs, j) {
			j--
		}
	}
	return true
}

func isalnum(bs []byte, i int) bool {
	if bs[i] >= 'A' && bs[i] <= 'Z' {
		bs[i] = bs[i] - 'A' + 'a'
		return true
	}

	if (bs[i] >= 'a' && bs[i] <= 'z') || (bs[i] >= '0' && bs[i] <= '9') {
		return true
	}

	return false
}

// 020
func countSubstrings(s string) int {
	dp := make([][]bool, len(s))

	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s); i++ {
			j := i + l - 1
			if j >= len(s) {
				break
			}

			if j-i < 3 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
			}
		}
	}

	cnt := 0
	for i := range dp {
		for j := range dp[i] {
			if dp[i][j] {
				cnt++
			}
		}
	}
	return cnt
}

// 021
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	second, first := dummy, head
	for i := 0; i < n; i++ {
		first = first.Next
	}

	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
	return dummy.Next
}

// 022
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			for {
				if slow == fast {
					return slow
				}
				slow = slow.Next
				fast = fast.Next
			}
		}
	}
	return nil
}

// 023
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil {
			curA = headB
		} else {
			curA = curA.Next
		}

		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA
}

// 024
func reverseList(head *ListNode) *ListNode {
	var cur, prev, temp *ListNode = head, nil, nil
	for cur != nil {
		temp = cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}

// 025
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]*ListNode, 0)
	stack2 := make([]*ListNode, 0)
	for l1 != nil || l2 != nil {
		if l1 != nil {
			stack1 = append(stack1, l1)
			l1 = l1.Next
		}

		if l2 != nil {
			stack2 = append(stack2, l2)
			l2 = l2.Next
		}
	}

	var node1, node2 *ListNode
	carry := 0
	for len(stack1) != 0 && len(stack2) != 0 {
		node1, node2 = stack1[len(stack1)-1], stack2[len(stack2)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]
		sum := node1.Val + node2.Val + carry
		node1.Val = sum % 10
		carry = sum / 10
	}

	if len(stack1) == 0 && len(stack2) != 0 {
		stack2[len(stack2)-1].Next = node1
		stack1 = append(stack1, stack2...)
	}

	for len(stack1) != 0 {
		node1 = stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		sum := node1.Val + carry
		node1.Val = sum % 10
		carry = sum / 10
	}

	if carry == 1 {
		node := &ListNode{1, node1}
		node1 = node
	}
	return node1
}

// 028
func flatten(root *Node) *Node {
	if root == nil {
		return root
	}

	next := make([]*Node, 0)
	cur := root
	for cur.Next != nil || cur.Child != nil || len(next) != 0 {
		if cur.Child != nil {
			if cur.Next != nil {
				next = append(next, cur.Next)
			}
			temp := cur.Child
			cur.Next = cur.Child
			cur.Child.Prev = cur
			cur.Child = nil
			cur = temp
		} else {
			if cur.Next != nil {
				cur = cur.Next
			} else {
				if len(next) > 0 {
					temp := next[len(next)-1]
					next = next[:len(next)-1]
					cur.Next = temp
					temp.Prev = cur
					cur = cur.Next
				}
			}
		}
	}
	return root
}

// 029
func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode = &Node{Val: x}
		aNode.Next = aNode
	} else {
		flag := false
		cur := aNode
		for {
			if (cur == aNode && flag) || (cur.Val < cur.Next.Val && x >= cur.Val && x <= cur.Next.Val) || (cur.Val > cur.Next.Val && (x <= cur.Next.Val || x >= cur.Val)) {
				node := &Node{Val: x}
				node.Next = cur.Next
				cur.Next = node
				break
			}

			if cur == aNode && !flag {
				flag = true
			}
			cur = cur.Next
		}
	}
	return aNode
}

// 032
func isAnagram(s string, t string) bool {
	cnt1, cnt2 := [26]int{}, [26]int{}
	for i := range s {
		cnt1[int(s[i]-'a')]++
	}

	for i := range t {
		cnt2[int(t[i]-'a')]++
	}

	return cnt1 == cnt2
}

// 033
func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int][]string)
	for i := range strs {
		temp := [26]int{}
		for j := range strs[i] {
			temp[int(strs[i][j]-'a')]++
		}
		hash[temp] = append(hash[temp], strs[i])
	}

	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}

// 034
func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)
	for i := range order {
		orderMap[order[i]] = i
	}

	for i := 0; i < len(words)-1; i++ {
		flag := false
		for j := 0; j < len(words[i]) && j < len(words[i+1]); j++ {
			if orderMap[words[i][j]] < orderMap[words[i+1][j]] {
				flag = true
				break
			}

			if orderMap[words[i][j]] > orderMap[words[i+1][j]] {
				return false
			}
		}

		if !flag && len(words[i]) > len(words[i+1]) {
			return false
		}
	}
	return true
}

// 036
func evalRPN(tokens []string) int {
	ans := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+":
			ans[len(ans)-2] = ans[len(ans)-2] + ans[len(ans)-1]
			ans = ans[:len(ans)-1]
		case "-":
			ans[len(ans)-2] = ans[len(ans)-2] - ans[len(ans)-1]
			ans = ans[:len(ans)-1]
		case "*":
			ans[len(ans)-2] = ans[len(ans)-2] * ans[len(ans)-1]
			ans = ans[:len(ans)-1]
		case "/":
			ans[len(ans)-2] = ans[len(ans)-2] / ans[len(ans)-1]
			ans = ans[:len(ans)-1]
		default:
			num, _ := strconv.Atoi(v)
			ans = append(ans, num)
		}
	}
	return ans[0]
}

// 037
func asteroidCollision(asteroids []int) []int {
	left := make([]int, 0)
	ans := make([]int, 0)
	for i := range asteroids {
		if asteroids[i] > 0 {
			left = append(left, asteroids[i])
		} else {
			temp := asteroids[i]
			for temp < 0 && len(left) != 0 {
				last := left[len(left)-1]
				left = left[:len(left)-1]
				if temp+last > 0 {
					temp = last
				} else if temp+last == 0 {
					temp = 0
				}
			}

			if temp < 0 {
				ans = append(ans, temp)
			} else if temp > 0 {
				left = append(left, temp)
			}
		}
	}

	ans = append(ans, left...)
	return ans
}

// 038
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	stack := []int{0}
	for i := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}
	return ans
}

// 098
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[n-1]
}

// 099
func minPathSum(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		dp[i] = grid[0][i] + dp[i-1]
	}

	for i := 1; i < len(grid); i++ {
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[len(dp)-1]
}

// 100
func minimumTotal(triangle [][]int) int {
	dp := make([]int, len(triangle))
	dp[len(dp)-1] = triangle[0][0]
	for i := 0; i < len(triangle)-1; i++ {
		dp[i] = 10000
	}

	for i := 1; i < len(triangle); i++ {
		for j := i + 1; j > 1; j-- {
			dp[len(dp)-j] = min(dp[len(dp)-j], dp[len(dp)-j+1]) + triangle[i][i+1-j]
		}
		dp[len(dp)-1] = dp[len(dp)-1] + triangle[i][i]
	}

	ans := dp[len(dp)-1]
	for i := len(dp) - 2; i >= 0; i-- {
		ans = min(ans, dp[i])
	}
	return ans
}

