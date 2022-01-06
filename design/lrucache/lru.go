package lrucache

type LRUCache struct {
	Cap  int
	Map  map[int]*Node
	Head *Node
	Tail *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	node := &Node{0, 0, nil, nil}
	return LRUCache{
		Cap:  capacity,
		Map:  make(map[int]*Node),
		Head: node,
		Tail: node,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.Map[key]
	if !exist {
		return -1
	}
	if node != this.Tail {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		node.Prev = this.Tail
		this.Tail.Next = node
		node.Next = nil
		this.Tail = node
	}
	return node.Val
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.Map[key]
	if exist {
		node.Val = value
		if node != this.Tail {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			node.Prev = this.Tail
			this.Tail.Next = node
			node.Next = nil
			this.Tail = node
		}
	} else {
		node = &Node{key, value, nil, nil}
		if len(this.Map) == this.Cap {
			delete(this.Map, this.Head.Next.Key)
			if this.Cap != 1 {
				this.Head.Next.Next.Prev = this.Head
				this.Head.Next = this.Head.Next.Next
			} else {
				this.Head.Next = nil
				this.Tail = this.Head
			}
		}
		this.Map[key] = node
		this.Tail.Next = node
		node.Prev = this.Tail
		this.Tail = node
	}
}
