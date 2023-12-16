package auto_complete_feature

const alphabetSize = 26 + 1

func Solution(q string, dict []string) []string {
	if q == "" {
		return nil
	}

	root := newTrieNode()
	for _, word := range dict {
		root.insert(word)
	}

	return root.autoComplete(q)
}

type trieNode struct {
	childs []*trieNode
	isWord bool
	isLeaf bool
}

func newTrieNode() *trieNode {
	return &trieNode{
		childs: make([]*trieNode, alphabetSize),
	}
}

func (node *trieNode) insert(s string) {
	if len(s) == 0 {
		node.isLeaf = true
		node.isWord = true
		return
	}

	index := ctoi(rune(s[0]))
	child := node.childs[index]
	if child == nil {
		child = newTrieNode()
		node.childs[index] = child
	}

	child.isLeaf = false
	child.insert(s[1:])
}

func (node *trieNode) autoComplete(s string) []string {
	next := node
	for _, r := range s {
		next = next.childs[ctoi(r)]
		if next == nil {
			return nil
		}
	}

	if next == nil {
		return nil
	}

	return next.append(s)
}

func (node *trieNode) append(s string) []string {
	var out []string
	if node.isWord {
		out = append(out, s)
	}

	for i, child := range node.childs {
		if child == nil {
			continue
		}

		prefix := s + string(itoc(rune(i)))

		if child.isLeaf {
			out = append(out, prefix)
			continue
		}

		for _, word := range child.append(prefix) {
			out = append(out, word)
		}
	}

	return out
}

func itoc(i rune) rune {
	switch i {
	case 26:
		return ' '
	default:
		return i + 'a'
	}
}

func ctoi(c rune) rune {
	switch c {
	case ' ':
		return 26
	default:
		return c - 'a'
	}
}
