
type PrefixTree struct {
	isWord   bool
	children map[rune]*PrefixTree
}

func Constructor() PrefixTree {
	return PrefixTree{
		isWord:   false,
		children: make(map[rune]*PrefixTree),
	}
}

func (this *PrefixTree) Insert(word string) {

	curr := this
	for _, c := range word {
		if node, ok := curr.children[c]; ok {
			curr = node
		} else {
			newNode := Constructor()
			curr.children[c] = &newNode
			curr = curr.children[c]
		}
	}
	curr.isWord = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this
	for _, c := range word {
		if node, ok := curr.children[c]; ok {
			curr = node
		} else {
			return false
		}
	}
	return curr.isWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this
	for _, c := range prefix {
		if node, ok := curr.children[c]; ok {
			curr = node
		} else {
			return false
		}
	}

	return true
}
