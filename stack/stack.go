package stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res
}

// 二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Right
	}
	return res
}

// 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)

	var cur, prev *TreeNode = root, nil
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		if cur.Right == nil || cur.Right == prev {
			res = append(res, cur.Val)
			stack = stack[:len(stack)-1]
			prev = cur
			cur = nil
		} else {
			cur = cur.Right
		}
	}
	return res
}

func flatten(root *TreeNode) {
}

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

func calculate(s string) int {
	number, operator := make([]int, 0), byte('+')
	num := 0
	for i := 0; i < len(s); i++ {
		if s[i] <= '9' && s[i] >= '0' {
			num = 10*num + int(s[i]-'0')
			if i == len(s)-1 || s[i+1] < '0' || s[i+1] > '9' {
				switch operator {
				case '+':
					number = append(number, num)
				case '-':
					number = append(number, -num)
				case '*':
					number[len(number)-1] = number[len(number)-1] * num
				case '/':
					number[len(number)-1] = number[len(number)-1] / num
				}
				num = 0
			}
		}

		if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
			operator = s[i]
		}
	}

	sum := 0
	for _, v := range number {
		sum += v
	}
	return sum
}