package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	second, first := dummy, head
	for i := 0; i < right-left; i++ {
		first = first.Next
	}

	for i := 0; i < left-1; i++ {
		second = second.Next
		first = first.Next
	}

	for second != nil && second != first {
		temp := second.Next
		second.Next = second.Next.Next
		temp.Next = first.Next
		first.Next = temp
	}

	return dummy.Next
}

// 快慢指针，快指针步长为慢指针的两倍
// 首次相遇，设慢指针走过n步，则快指针走过2n步，n即为环的长度
// 设慢指针从head走到环的入口节点为x步，则慢指针在环中走过n-x步
// 快指针重新指向head，步长与慢指针一致，同时走过x步后在环的入口节点相遇
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

func swapNodes(head *ListNode, k int) *ListNode {
	second, first := head, head
	for i := 1; i < k; i++ {
		second = second.Next
	}

	temp := second

	for temp.Next != nil {
		first = first.Next
		temp = temp.Next
	}

	second.Val, first.Val = first.Val, second.Val

	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{0, l1}
	cur1, cur2 := dummy, l2
	carry := 0
	for cur1.Next != nil && cur2 != nil {
		sum := cur1.Next.Val + cur2.Val + carry
		cur1.Next.Val = sum % 10
		carry = sum / 10

		cur1 = cur1.Next
		cur2 = cur2.Next
	}

	if cur1.Next == nil && cur2 != nil {
		cur1.Next = cur2
	}

	for cur1.Next != nil {
		sum := cur1.Next.Val + carry
		cur1.Next.Val = sum % 10
		carry = sum / 10
		cur1 = cur1.Next
	}

	if carry == 1 {
		cur1.Next = &ListNode{1, nil}
	}

	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next.Next
		cur.Next.Next = temp.Next
		temp.Next = cur.Next
		cur.Next = temp

		cur = cur.Next.Next
	}

	return dummy.Next
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 0
	cur := head
	for cur != nil {
		l++
		cur = cur.Next
	}

	if l == 0 || k%l == 0 {
		return head
	}

	second, first := head, head
	for i := 0; i < k%l; i++ {
		first = first.Next
	}
	for first.Next != nil {
		second = second.Next
		first = first.Next
	}
	first.Next = head
	res := second.Next
	second.Next = nil
	return res
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == temp {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// 从head开始找到第一个val大于等于x的节点
// 将该节点后面val小于x的节点移动至该节点之前
func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{0, head}
	second := dummy
	for second.Next != nil {
		if second.Next.Val >= x {
			break
		}
		second = second.Next
	}

	first := second
	for first.Next != nil {
		if first.Next.Val < x {
			temp := first.Next
			first.Next = first.Next.Next
			temp.Next = second.Next
			second.Next = temp
			second = second.Next
		} else {
			first = first.Next
		}
	}
	return dummy.Next
}
