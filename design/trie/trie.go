package trie

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, v := range word {
		idx := v - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &Trie{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for _, v := range prefix {
		idx := v - 'a'
		if node.children[idx] == nil {
			return nil
		}
		node = node.children[idx]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
