package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	ans := make([]int, 0)
	for len(queue) != 0 {
		max := queue[0].Val
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i].Val > max {
				max = queue[i].Val
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		ans = append(ans, max)
		queue = queue[l:]
	}
	return ans
}

func findBottomLeftValue(root *TreeNode) int {
	ans := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		ans = queue[0].Val
		l := len(queue)
		for i := 0; i < l; i++ {

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return ans
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	ans := make([]int, 0)
	for len(queue) != 0 {
		ans = append(ans, queue[len(queue)-1].Val)
		l := len(queue)
		for i := 0; i < l; i++ {

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[l:]
	}
	return ans
}

func pruneTree(root *TreeNode) *TreeNode {
	if !contain1Tree(root) {
		root = nil
	}
	return root
}

func contain1Tree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	left := contain1Tree(root.Left)
	right := contain1Tree(root.Right)
	if !left {
		root.Left = nil
	}

	if !right {
		root.Right = nil
	}

	return root.Val == 1 || left || right
}

func sumNumbers(root *TreeNode) int {
	ans, temp := 0, 0
	stack := make([]*TreeNode, 0)
	var cur, prev *TreeNode = root, nil
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			temp = temp*10 + cur.Val
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		if cur.Left == nil && cur.Right == nil {
			ans += temp
		}
		if cur.Right == nil || cur.Right == prev {
			stack = stack[:len(stack)-1]
			temp /= 10
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return ans
}
