package sword

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
