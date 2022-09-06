package lc3

type Trie struct {
	children []*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	node := trie
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
		}
		node = node.children[c-'a']
	}
	node.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	node := trie.searchPrefix(word)
	return node != nil && node.isEnd
}

func (trie *Trie) StartsWith(prefix string) bool {
	return trie.searchPrefix(prefix) != nil
}

func (trie *Trie) searchPrefix(prefix string) *Trie {
	node := trie
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return nil
		}
		node = node.children[c-'a']
	}
	return node
}
