package backtrack

import (
	"fmt"
	"sort"
	"strconv"
)

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0, len(nums))
	used := make(map[int]struct{})
	dfsPermute(nums, &ans, &temp, used)
	return ans
}

func dfsPermute(nums []int, ans *[][]int, temp *[]int, used map[int]struct{}) {
	if len(*temp) == len(nums) {
		copy := append(make([]int, 0, len(*temp)), *temp...)
		*ans = append(*ans, copy)
		return
	}

	for i := range nums {
		if _, ok := used[nums[i]]; !ok {
			*temp = append(*temp, nums[i])
			used[nums[i]] = struct{}{}
			dfsPermute(nums, ans, temp, used)
			*temp = (*temp)[:len(*temp)-1]
			delete(used, nums[i])
		}
	}
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	used := make([]bool, len(nums))
	temp := make([]int, 0, len(nums))
	dfsPermuteUnique(nums, used, &temp, &ans)
	return ans
}

func dfsPermuteUnique(nums []int, used []bool, temp *[]int, ans *[][]int) {
	if len(*temp) == len(nums) {
		copy := append(make([]int, 0), (*temp)...)
		*ans = append(*ans, copy)
		return
	}

	for i := range nums {
		if used[i] || (i > 0 && nums[i] == nums[i-1] && !used[i-1]) {
			continue
		}

		*temp = append(*temp, nums[i])
		used[i] = true
		dfsPermuteUnique(nums, used, temp, ans)
		used[i] = false
		*temp = (*temp)[:len(*temp)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0)
	sum := 0
	dfsCombinationSum(candidates, &ans, &temp, target, &sum)
	return ans
}

func dfsCombinationSum(candidates []int, ans *[][]int, temp *[]int, target int, sum *int) {
	if *sum > target {
		return
	}

	if *sum == target {
		copy := append(make([]int, 0), (*temp)...)
		*ans = append(*ans, copy)
		return
	}

	for i := range candidates {

		*temp = append(*temp, candidates[i])
		*sum += candidates[i]
		dfsCombinationSum(candidates, ans, temp, target, sum)
		*temp = (*temp)[:len(*temp)-1]
		*sum -= candidates[i]
	}
}

// 37. 解数独
func SolveSudoku(board [][]byte) {
	rows := [9][9]bool{}
	columns := [9][9]bool{}
	areas := [9][9]bool{}
	blanks := make([]int, 0)
	for i := range board {
		for j := range board[0] {
			if board[i][j] == '.' {
				blanks = append(blanks, i*9+j)
			} else {
				rows[i][board[i][j]-'1'] = true
				columns[j][board[i][j]-'1'] = true
				areas[i/3*3+j/3][board[i][j]-'1'] = true
			}
		}
	}

	dict := []byte("123456789")
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == len(blanks) {
			return true
		}

		i, j := blanks[idx]/9, blanks[idx]%9
		for k := range dict {
			if rows[i][dict[k]-'1'] || columns[j][dict[k]-'1'] || areas[i/3*3+j/3][dict[k]-'1'] {
				continue
			}
			board[i][j] = dict[k]
			rows[i][dict[k]-'1'] = true
			columns[j][dict[k]-'1'] = true
			areas[i/3*3+j/3][dict[k]-'1'] = true
			if dfs(idx + 1) {
				return true
			}
			rows[i][dict[k]-'1'] = false
			columns[j][dict[k]-'1'] = false
			areas[i/3*3+j/3][dict[k]-'1'] = false
		}
		return false
	}

	dfs(0)
}

// 51. N 皇后
func SolveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	ans := make([][]string, 0)
	columns := make([]bool, n)
	diagonals := make(map[int]bool) // 斜线

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			temp := make([]string, 0, n)
			for i := range board {
				temp = append(temp, string(append([]byte{}, board[i]...)))
			}
			ans = append(ans, temp)
			return
		}

		for i := 0; i < n; i++ {
			if columns[i] || diagonals[i+row+n] || diagonals[i-row] {
				continue
			}

			board[row][i] = 'Q'
			columns[i] = true
			diagonals[i+row+n] = true
			diagonals[i-row] = true
			dfs(row + 1)
			board[row][i] = '.'
			columns[i] = false
			diagonals[i+row+n] = false
			diagonals[i-row] = false
		}

	}
	dfs(0)
	return ans
}

// 52. N皇后 II
func totalNQueens(n int) int {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	ans := 0
	columns := make([]bool, n)
	diagonals := make(map[int]bool) // 斜线

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			ans++
			return
		}

		for i := 0; i < n; i++ {
			if columns[i] || diagonals[i+row+n] || diagonals[i-row] {
				continue
			}

			board[row][i] = 'Q'
			columns[i] = true
			diagonals[i+row+n] = true
			diagonals[i-row] = true
			dfs(row + 1)
			board[row][i] = '.'
			columns[i] = false
			diagonals[i+row+n] = false
			diagonals[i-row] = false
		}

	}
	dfs(0)
	return ans
}

// 93. 复原 IP 地址
func restoreIpAddresses(s string) []string {
	idxs := []int{}
	ans := make([]string, 0)

	var dfs func(step, start int)
	dfs = func(step, start int) {
		if step == 4 || start == len(s) {
			if start == len(s) && step == 4 {
				ans = append(ans, fmt.Sprintf("%s.%s.%s.%s", s[:idxs[0]], s[idxs[0]:idxs[1]], s[idxs[1]:idxs[2]], s[idxs[2]:]))
			}
			return
		}

		l := 3
		if s[start] == '0' {
			l = 1
		}

		for i := 1; i <= l; i++ {
			if start+i > len(s) {
				break
			}

			n, _ := strconv.Atoi(s[start : start+i])
			if n > 255 {
				continue
			}
			idxs = append(idxs, start+i)
			dfs(step+1, start+i)
			idxs = idxs[:step]
		}
	}
	dfs(0, 0)

	return ans
}

// 剑指 Offer 38. 字符串的排列
func permutation(s string) []string {
	bs := []byte(s)

	sort.Slice(bs, func(i, j int) bool {
		return bs[i] < bs[j]
	})

	ans := make([]string, 0)
	bytes := make([]byte, 0, len(s))
	used := make([]bool, len(s))

	var dfs func(step int)
	dfs = func(step int) {
		if step == len(s) {
			ans = append(ans, string(bytes))
			return
		}

		for i := range bs {
			if used[i] || (i > 0 && bs[i] == bs[i-1] && !used[i-1]) {
				continue
			}
			bytes = append(bytes, bs[i])
			used[i] = true
			dfs(step + 1)
			bytes = bytes[:step]
			used[i] = false
		}
	}
	dfs(0)
	return ans
}

// 78. 子集
func subsets(nums []int) [][]int {

	ans := make([][]int, 0)
	arr := make([]int, 0)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		dfs(cur + 1)
		arr = append(arr, nums[cur])
		dfs(cur + 1)
		arr = arr[:len(arr)-1]
	}
	dfs(0)
	return ans
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	ans := make([][]int, 0)
	arr := make([]int, 0)
	used := make([]bool, len(candidates))

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

		if idx > 0 && candidates[idx] == candidates[idx-1] && !used[idx-1] {
			return
		}

		if target-candidates[idx] >= 0 {
			arr = append(arr, candidates[idx])
			used[idx] = true
			dfs(target-candidates[idx], idx+1)
			arr = arr[:len(arr)-1]
			used[idx] = false
		}
	}
	dfs(target, 0)
	return ans
}

// 494. 目标和
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	sum := 0

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == len(nums) {
			if sum == target {
				ans++
			}
			return
		}

		sum -= nums[idx]
		dfs(idx + 1)

		sum += 2 * nums[idx]
		dfs(idx + 1)
		sum -= nums[idx]
	}

	dfs(0)
	return ans
}

// 526. 优美的排列
func countArrangement(n int) int {
	used := make([]bool, n)

	ans := 0

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == n+1 {
			ans++
			return
		}

		for i := 1; i <= n; i++ {
			if used[i-1] {
				continue
			}
			if i%idx == 0 || idx%i == 0 {
				used[i-1] = true
				dfs(idx + 1)
				used[i-1] = false
			}
		}
	}
	dfs(1)
	return ans
}

// 77. 组合
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	arr := make([]int, 0)
	var dfs func(num, cnt int)
	dfs = func(num, cnt int) {
		if cnt == 0 {
			ans = append(ans, append([]int{}, arr...))
			return
		}

		if num == n+1 {
			return
		}

		dfs(num+1, cnt)

		arr = append(arr, num)
		dfs(num+1, cnt-1)
		arr = arr[:len(arr)-1]
	}
	dfs(1, k)
	return ans
}

func generateParenthesis(n int) []string {

}