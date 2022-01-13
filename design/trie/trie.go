package trie

type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
}

func (this *Trie) Search(word string) bool {
}

func (this *Trie) StartsWith(prefix string) bool {
}
