package backtrack

import "sort"

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

func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	temp := make([]int, 0, k)
	dfsCombine(n, k, &ans, &temp)
	return ans
}

func dfsCombine(n int, k int, ans *[][]int, temp *[]int) {
	if len(*temp) == k {
		copy := append(make([]int, 0, k), *temp...)
		*ans = append(*ans, copy)
		return
	}

	for i := 1; i <= n; i++ {
		if len(*temp) == 0 || i > (*temp)[len(*temp)-1] {
			*temp = append(*temp, i)
			dfsCombine(n, k, ans, temp)
			*temp = (*temp)[:len(*temp)-1]
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
