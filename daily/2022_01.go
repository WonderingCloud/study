package daily

// 2022-01-01
func construct2DArray(original []int, m int, n int) [][]int {
	if m*n != len(original) {
		return nil
	}

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = original[i*n : (i+1)*n]
	}
	return ans
}

// 2022-01-02
func lastRemaining(n int) int {
	ans := 1
	k, cnt, step := 0, n, 1
	for cnt > 1 {
		if k%2 == 0 { // 正向
			ans += step
		} else { // 反向
			if cnt%2 == 1 {
				ans += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return ans
}

// 2022-01-03
func dayOfTheWeek(day, month, year int) string {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	days := 0
	days += 365*(year-1971) + (year-1969)/4
	for _, d := range monthDays[:month-1] {
		days += d
	}
	if month >= 3 && (year%400 == 0 || year%4 == 0 && year%100 != 0) {
		days++
	}
	days += day
	return week[(days+3)%7]
}

// 2022-01-04
const (
	draw     = 0
	mouseWin = 1
	catWin   = 2
)

func catMouseGame(graph [][]int) int {
	n := len(graph)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n*2)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var getResult, getNextResult func(int, int, int) int
	getResult = func(mouse, cat, turns int) int {
		if turns == n*2 {
			return draw
		}
		res := dp[mouse][cat][turns]
		if res != -1 {
			return res
		}
		if mouse == 0 {
			res = mouseWin
		} else if cat == mouse {
			res = catWin
		} else {
			res = getNextResult(mouse, cat, turns)
		}
		dp[mouse][cat][turns] = res
		return res
	}
	getNextResult = func(mouse, cat, turns int) int {
		curMove := mouse
		if turns%2 == 1 {
			curMove = cat
		}
		defaultRes := mouseWin
		if curMove == mouse {
			defaultRes = catWin
		}
		res := defaultRes
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				continue
			}
			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}
			nextRes := getResult(nextMouse, nextCat, turns+1)
			if nextRes != defaultRes {
				res = nextRes
				if res != draw {
					break
				}
			}
		}
		return res
	}
	return getResult(1, 2, 0)
}

// 2022-01-05
func modifyString(s string) string {
	ans := []byte(s)
	for i := 0; i < len(ans); i++ {
		if ans[i] != '?' {
			continue
		}
		char := byte('a')
		for (i > 0 && char == ans[i-1]) || (i < len(s)-1 && char == ans[i+1]) {
			char++
		}
		ans[i] = char
	}
	return string(ans)
}

// 2022-01-06
func simplifyPath(path string) string {
	stack := make([]string, 0)
	j, s := 0, ""
	for i := range path {
		if path[i] == '/' {
			if i > 0 && path[i-1] != '/' {
				s = path[j:i]
			}
			j = i + 1
		}

		if i == len(path)-1 && path[i] != '/' {
			s = path[j:]
		}

		if len(s) != 0 {
			if s == ".." && len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else if s != ".." && s != "." {
				stack = append(stack, s)
			}
			s = ""
		}
	}

	ans := ""
	for i := range stack {
		ans = ans + "/" + stack[i]
	}

	if len(ans) == 0 {
		ans = "/"
	}
	return ans
}

// 2022-01-07
func maxDepth(s string) int {
	ans, cnt := 0, 0
	for i := range s {
		if s[i] == '(' {
			cnt++
			ans = max(ans, cnt)
		} else if s[i] == ')' {
			cnt--
		}
	}
	return ans
}

// 2022-01-08
func grayCode(n int) []int {
	ans := make([]int, 0)
	used := make(map[int]struct{})

	dfs := func(num int) {}
	dfs = func(num int) {
		used[num] = struct{}{}
		ans = append(ans, num)
		for i := 0; i < n; i++ {
			next := num ^ 1<<i
			if _, exist := used[next]; !exist {
				dfs(next)
			}
		}
	}

	dfs(0)
	return ans
}

// 2022-01-09
func slowestKey(releaseTimes []int, keysPressed string) byte {
	ans, maxTime, t := keysPressed[0], releaseTimes[0], 0
	for i := 1; i < len(releaseTimes); i++ {
		t = releaseTimes[i] - releaseTimes[i-1]
		if t > maxTime {
			maxTime = t
			ans = keysPressed[i]
		} else if t == maxTime && keysPressed[i] > ans {
			ans = keysPressed[i]
		}
	}
	return ans
}
