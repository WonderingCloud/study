package lrucache

type Node struct {
	Key  uint16
	Val  uint16
	Next *Node
	Prev *Node
}

type LRUCache struct {
	Cap  int16
	Map  map[uint16]*Node
	Head *Node
	Tail *Node
}

func Constructor(capacity int) LRUCache {
	node := new(Node)
	return LRUCache{
		Map:  make(map[uint16]*Node),
		Cap:  int16(capacity),
		Head: node,
		Tail: node,
	}
}

func (this *LRUCache) Get(key int) int {
	node, exist := this.Map[uint16(key)]
	if !exist {
		return -1
	}
	if node != this.Tail {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		this.Tail.Next = node
		node.Prev = this.Tail
		node.Next = nil
		this.Tail = node
	}
	return int(node.Val)
}

func (this *LRUCache) Put(key int, value int) {
	node, exist := this.Map[uint16(key)]
	if exist {
		node.Val = uint16(value)
		if node != this.Tail {
			node.Prev.Next = node.Next
			node.Next.Prev = node.Prev
			this.Tail.Next = node
			node.Prev = this.Tail
			node.Next = nil
			this.Tail = node
		}
		return
	}

	if len(this.Map) == int(this.Cap) {
		delete(this.Map, this.Head.Next.Key)
		if this.Head.Next.Next != nil {
			this.Head.Next.Next.Prev = this.Head
			this.Head.Next = this.Head.Next.Next
		} else {
			this.Head.Next = nil
			this.Tail = this.Head
		}
	}
	node = &Node{uint16(key), uint16(value), nil, this.Tail}
	this.Map[uint16(key)] = node
	this.Tail.Next = node
	this.Tail = node
}
