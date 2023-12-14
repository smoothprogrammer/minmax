package word_break_problem

import "strings"

func Solution(s string, dict []string) []string {
	if s == "" {
		return nil
	}

	root := newTrieNode()
	for _, word := range dict {
		root.insert(word)
	}

	var out []string
	queue := []wordBreak{{s: s}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.s == "" {
			out = append(out, strings.Join(curr.word, " "))
			continue
		}

		words := root.words(curr.s)
		for _, word := range words {
			queue = append(queue, wordBreak{
				word: append(curr.word, word),
				s:    curr.s[len(word):],
			})
		}
	}

	return out
}

type wordBreak struct {
	word []string // word break
	s    string   // unprocessed word
}

type trieNode struct {
	childs []*trieNode
	isWord bool
}

func newTrieNode() *trieNode {
	return &trieNode{
		childs: make([]*trieNode, 26),
	}
}

func (node *trieNode) insert(s string) {
	if len(s) == 0 {
		node.isWord = true
		return
	}

	index := s[0] - 'a'
	child := node.childs[index]
	if child == nil {
		child = newTrieNode()
		node.childs[index] = child
	}

	child.insert(s[1:])
}

func (node *trieNode) words(s string) []string {
	var out []string
	var buf strings.Builder
	next := node
	for len(s) > 0 {
		c := s[0]
		s = s[1:]
		index := c - 'a'
		next = next.childs[index]
		if next == nil {
			break
		}
		buf.WriteByte(c)
		if next.isWord {
			out = append(out, buf.String())
		}
	}
	return out
}
